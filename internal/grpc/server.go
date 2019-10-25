package grpc

import (
	"context"
	"encoding/json"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/omerkaya1/go-calendar/internal/domain/services"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
	queue "github.com/omerkaya1/go-calendar/internal/message-queue"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type GoCalendarServer struct {
	Cfg          *conf.Config
	Logger       *zap.Logger
	EventService *services.EventService
	Queue        *queue.MessageQueue
}

func NewServer(cfg *conf.Config, log *zap.Logger, es *services.EventService, q *queue.MessageQueue) *GoCalendarServer {
	return &GoCalendarServer{
		Cfg:          cfg,
		Logger:       log,
		EventService: es,
		Queue:        q,
	}
}

func (s *GoCalendarServer) Run() {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		s.Logger.Sugar().Errorf("%s", err)
	}

	gca.RegisterGoCalendarServerServer(server, s)

	// Producer both scans and enqueues upcoming events into a message queue
	go s.Producer()
	// Receiver emulates the process of receiving messages from the message queue
	go s.Receiver()

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", server.Serve(l))
}

func (s *GoCalendarServer) Producer() {
	s.Logger.Sugar().Info("Producer routine has started")
	ch, err := s.Queue.Conn.Channel()
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	q, err := ch.QueueDeclare(
		"events",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	// Create a ticker to trigger the the scan process and do the DB query
	wakeyWakey := time.NewTicker(s.Queue.QueryInterval)
	for range wakeyWakey.C {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		events, err := s.EventService.GetUpcomingEvents(ctx)
		if err != nil {
			s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		}
		if events == nil {
			break
		}
		for _, e := range events {
			body, err := json.Marshal(e)
			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			if err != nil {
				s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
			}
		}
	}
}

func (s *GoCalendarServer) Receiver() {
	s.Logger.Sugar().Info("Receiver routine has started")
	conn, err := queue.NewMessageQueue(s.Cfg.Queue)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	ch, err := conn.Conn.Channel()
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	q, err := ch.QueueDeclare(
		"events",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}

	for d := range msgs {
		e := models.Event{}
		err := json.Unmarshal(d.Body, &e)
		if err != nil {
			s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
			continue
		}
		s.Logger.Sugar().Infof("message received: %v", e)
	}
	s.Logger.Sugar().Info("The message queue channel was closed")
	return
}
