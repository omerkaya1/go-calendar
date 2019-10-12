package models

import (
	"github.com/satori/go.uuid"
	"time"
)

// Event
type Event struct {
	EventId   uuid.UUID  `json:"event_id"`
	UserName  string     `json:"user_name"`
	EventName string     `json:"event_name"`
	Note      string     `json:"note"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}

// NewEvent .
func NewEvent(user, event, note string, start, end *time.Time) *Event {
	return &Event{
		EventId:   uuid.NewV4(),
		UserName:  user,
		EventName: event,
		Note:      note,
		StartTime: start,
		EndTime:   end,
	}
}

// ComposeEvent .
func ComposeEvent(old Event, new *Event) *Event {
	retEvent := &Event{
		EventId:   old.EventId,
		UserName:  old.UserName,
		EventName: old.EventName,
		Note:      old.Note,
		StartTime: old.StartTime,
		EndTime:   old.EndTime,
	}
	if new.EventName != "" {
		retEvent.EventName = new.EventName
	}
	if new.StartTime != nil {
		retEvent.StartTime = new.StartTime
	}
	if new.EndTime != nil {
		retEvent.EndTime = new.EndTime
	}
	return retEvent
}
