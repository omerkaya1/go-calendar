package rws

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/omerkaya1/go-calendar/internal/domain/services"
	queue "github.com/omerkaya1/go-calendar/internal/message-queue"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"net/http"
	"path"
	"time"
)

type Server struct {
	Cfg          *conf.Config
	Logger       *zap.Logger
	EventService *services.EventService
	Queue        *queue.MessageQueue
}

func NewServer(cfg *conf.Config, log *zap.Logger, es *services.EventService, q *queue.MessageQueue) *Server {
	return &Server{
		Cfg:          cfg,
		Logger:       log,
		EventService: es,
		Queue:        q,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL), s.CreateEvent).Methods(http.MethodPost)
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL, "/{id:[-A-Z0-9a-z]+}"), s.GetEvent).Methods(http.MethodGet)
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL), s.UpdateEvent).Methods(http.MethodPut)
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL, "/{id:[-A-Z0-9a-z]+}"), s.DeleteEvent).Methods(http.MethodDelete)

	// Producer both scans and enqueues upcoming events into a message queue
	go s.Producer()
	// Receiver emulates the process of receiving messages from the message queue
	go s.Receiver()

	s.Logger.Sugar().Infof("Server initialised on address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", http.ListenAndServe(s.Cfg.Host+":"+s.Cfg.Port, r))
}

func (s *Server) Producer() {
	s.Logger.Sugar().Info("Producer routine has started")
	ch, err := s.Queue.Conn.Channel()
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	q, err := ch.QueueDeclare(
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
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

func (s *Server) Receiver() {
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
