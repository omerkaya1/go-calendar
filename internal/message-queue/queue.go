package message_queue

import (
	"context"
	"fmt"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/interfaces"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type MessageQueue struct {
	Conn          *amqp.Connection
	QueryInterval time.Duration
	db            interfaces.EventStorageProcessor
}

func NewMessageQueue(conf conf.QueueConf, db interfaces.EventStorageProcessor) (*MessageQueue, error) {
	if conf.Host == "" || conf.Port == "" || conf.User == "" || conf.Password == "" {
		return nil, errors.ErrBadQueueConfiguration
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port))
	if err != nil {
		return nil, err
	}
	var interval time.Duration
	if conf.Interval == "" {
		interval = 2 * time.Minute
	} else {
		i, err := time.ParseDuration(conf.Interval)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return &MessageQueue{Conn: conn, QueryInterval: interval, db: db}, nil
}

func (mq *MessageQueue) Produce() {
	log.Println("Producer routine has started")
	ch, err := mq.Conn.Channel()
	if err != nil {
		log.Printf("%s: %s", errors.ErrMQPrefix, err)
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
		log.Printf("%s: %s", errors.ErrMQPrefix, err)
		return
	}
	// Create a ticker to trigger the the scan process and do the DB query
	wakeyWakey := time.NewTicker(mq.QueryInterval)
	for range wakeyWakey.C {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		events, err := mq.db.GetUpcomingEvents(ctx)
		if err != nil {
			log.Printf("%s: %s", errors.ErrMQPrefix, err)
			continue
		}
		if events != nil {
			for _, e := range events {
				body := fmt.Sprintf("User: %s has '%s' from %s until %s",
					e.UserName, e.EventName, e.StartTime, e.EndTime)
				err := ch.Publish(
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

func (mq *MessageQueue) Receive() {
	log.Println("Receiver routine has started")
	ch, err := mq.Conn.Channel()
	if err != nil {
		log.Printf("%s: %s", errors.ErrMQPrefix, err)
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
		log.Printf("%s: %s", errors.ErrMQPrefix, err)
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
		log.Printf("%s: %s", errors.ErrMQPrefix, err)
		return
	}

	for d := range msgs {
		log.Printf("message received: %v", string(d.Body))
	}
	log.Println("The message queue channel was closed")
}
