package message_queue

import (
	"fmt"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/streadway/amqp"
	"time"
)

type MessageQueue struct {
	Conn          *amqp.Connection
	QueryInterval time.Duration
}

func NewMessageQueue(conf conf.QueueConf) (*MessageQueue, error) {
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
	return &MessageQueue{Conn: conn, QueryInterval: interval}, nil
}
