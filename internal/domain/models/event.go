package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Event struct {
	EventId   uuid.UUID
	UserName  string
	EventName string
	Note      string
	StartTime *time.Time
	EndTime   *time.Time
}
