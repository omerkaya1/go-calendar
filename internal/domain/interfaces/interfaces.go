package interfaces

import (
	"context"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/satori/go.uuid"
)

// EventStorageProcessor is an interface to communicate with the DB
type EventStorageProcessor interface {
	// GetOwnerEvents returns a slice of events that were created by the specified user
	GetUserEvents(context.Context, string) ([]models.Event, error)
	// GetEventByID returns the requested event, which id was specified by the callee
	GetEventByID(context.Context, uuid.UUID) (models.Event, error)
	// GetEventByName returns the requested event, which name was specified by the callee
	GetEventByName(context.Context, string) (models.Event, error)
	// UpdateEventByID updates an event stored in the DB by its internal ID
	UpdateEventByID(context.Context, uuid.UUID, *models.Event) (uuid.UUID, error)
	// UpdateEventByID updates an event stored in the DB by its name
	UpdateEventByName(context.Context, string, *models.Event) (uuid.UUID, error)
	// CreateEvent creates a new event
	CreateEvent(context.Context, *models.Event) (uuid.UUID, error)
	// DeleteEventById deletes an event by its internal ID
	DeleteEventById(context.Context, uuid.UUID) error
	// DeleteAllUserEvents deletes all the events stored in the DB under a specified user
	DeleteAllUserEvents(context.Context, string) error
}
