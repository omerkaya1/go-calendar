package mq

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/interfaces"
	"github.com/omerkaya1/go-calendar/internal/watcher/errors"
	"github.com/streadway/amqp"
)

// MessageQueue
type MessageQueue interface {
	ProduceMessages() error
	EmmitMessages() error
}

// EventMQProducer .
// TODO: in order to communicate with the db we only need a specific set of methods strictly limited to reading!
// 		 Therefore, we should define a smaller interface to both satisfy our needs and comply with the
type RabbitMQService struct {
	Conn  *amqp.Connection
	db    interfaces.EventStorageProcessor
	conf  *config.Config
	count prometheus.Counter
}

// NewEventMQProducer .
func NewRabbitMQService(conf *config.Config, db interfaces.EventStorageProcessor, mc prometheus.Counter) (MessageQueue, error) {
	if conf.Queue.Host == "" || conf.Queue.Port == "" || conf.Queue.User == "" || conf.Queue.Password == "" || conf.Queue.Name == "" {
		return nil, errors.ErrBadQueueConfiguration
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Queue.User, conf.Queue.Password, conf.Queue.Host, conf.Queue.Port))
	if err != nil {
		return nil, err
	}

	return &RabbitMQService{Conn: conn, conf: conf, db: db, count: mc}, nil
}

// ProduceMessages .
func (rms *RabbitMQService) ProduceMessages() error {
	ch, err := rms.Conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		rms.conf.Queue.Name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	interval, err := time.ParseDuration(rms.conf.Queue.Interval)
	if err != nil {
		return err
	}

	// Handle interrupt signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	// Create a ticker to trigger the the scan process and do the DB query
	tickTockBoom := time.NewTicker(interval)

MQ:
	for {
		select {
		case <-stopChan:
			ch.Close()
			rms.Conn.Close()
			log.Println("Exit the programme.")
			break MQ
		case <-tickTockBoom.C:
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			events, err := rms.db.GetUpcomingEvents(ctx)
			if err != nil {
				log.Printf("%s: %s", errors.ErrMQPrefix, err)
				break
			}
			if events != nil {
				for _, e := range events {
					body := fmt.Sprintf("User: %s has '%s' from %s until %s",
						e.UserName, e.EventName, e.StartTime, e.EndTime)
					err = ch.Publish(
						"",
						q.Name,
						false,
						false,
						amqp.Publishing{
							ContentType: "application/json",
							Body:        []byte(body),
						})
					if err != nil {
						log.Printf("%s: %s", errors.ErrMQPrefix, err)
					}
				}
			}
		}
	}
	return nil
}

// EmmitMessages .
func (rms *RabbitMQService) EmmitMessages() error {
	ch, err := rms.Conn.Channel()
	if err != nil {
		return err
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
		return err
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
		return err
	}

	// Handle interrupt
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)

MQ:
	for {
		select {
		case <-exitChan:
			log.Println("Exit the programme.")
			ch.Close()
			rms.Conn.Close()
			break MQ
		case d, ok := <-msgs:
			if !ok {
				break MQ
			}
			log.Println(string(d.Body))
			rms.count.Inc()
		}
	}
	return nil
}

func (rms *RabbitMQService)
