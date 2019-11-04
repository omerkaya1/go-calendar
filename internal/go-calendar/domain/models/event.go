package models

import (
	"time"

	"github.com/satori/go.uuid"
)

// Event struct is the main internal representation of an event
type Event struct {
	EventID   uuid.UUID  `json:"event_id" db:"id"`
	UserName  string     `json:"user_name" db:"user_name"`
	EventName string     `json:"event_name" db:"title"`
	Note      string     `json:"note" db:"note"`
	StartTime *time.Time `json:"start_time" db:"start_time"`
	EndTime   *time.Time `json:"end_time" db:"end_time"`
}

// EventJSON struct is used for RWS Client-Server communications
type EventJSON struct {
	EventID   string `json:"event_id"`
	UserName  string `json:"user_name"`
	EventName string `json:"event_name"`
	Note      string `json:"note"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// NewEvent returns a new Event object to the callee
func NewEvent(user, event, note string, start, end *time.Time) *Event {
	return &Event{
		EventID:   uuid.NewV4(),
		UserName:  user,
		EventName: event,
		Note:      note,
		StartTime: start,
		EndTime:   end,
	}
}

// ComposeEvent method updates fields of an old event if necessary and returns it to the callee
func ComposeEvent(old Event, new *Event) *Event {
	retEvent := &Event{
		EventID:   old.EventID,
		UserName:  old.UserName,
		EventName: old.EventName,
		Note:      old.Note,
		StartTime: old.StartTime,
		EndTime:   old.EndTime,
	}
	if new.Note != "" {
		retEvent.Note = new.Note
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
