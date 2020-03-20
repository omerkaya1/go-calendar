package mq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/interfaces"
	"github.com/omerkaya1/go-calendar/internal/watcher/errors"
	"github.com/streadway/amqp"
)

// MessageQueue declares a general interface to operate message queues
type MessageQueue interface {
	ProduceMessages(context.Context) error
	EmmitMessages(context.Context) error
}

// RabbitMQService .
// TODO: in order to communicate with the db we only need a specific set of methods strictly limited to reading!
// 		 Therefore, we should define a smaller interface to both satisfy our needs and comply with the
type RabbitMQService struct {
	Conn  *amqp.Connection
	db    interfaces.EventStorageProcessor
	conf  config.QueueConf
	count prometheus.Counter
}

// NewRabbitMQService returns a new instance of MessageQueue interface
func NewRabbitMQService(conf config.QueueConf, db interfaces.EventStorageProcessor, mc prometheus.Counter) (MessageQueue, error) {
	if !conf.Verify() {
		return nil, errors.ErrBadQueueConfiguration
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port))
	if err != nil {
		return nil, err
	}

	return &RabbitMQService{Conn: conn, conf: conf, db: db, count: mc}, nil
}

// ProduceMessages queries the app's DB for upcoming events and enqueues them
func (rms *RabbitMQService) ProduceMessages(ctx context.Context) error {
	ch, err := rms.Conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		rms.conf.Name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	interval, err := time.ParseDuration(rms.conf.Interval)
	if err != nil {
		return err
	}
	// Create a ticker to trigger the the scan process and do the DB query
	tickTockBoom := time.NewTicker(interval)

MQ:
	for {
		select {
		case <-ctx.Done():
			ch.Close()
			rms.Conn.Close()
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

// EmmitMessages emits received messages to the stdout
func (rms *RabbitMQService) EmmitMessages(ctx context.Context) error {
	ch, err := rms.Conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		rms.conf.Name,
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
	// Main work cycle
MQ:
	for {
		select {
		case <-ctx.Done():
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
