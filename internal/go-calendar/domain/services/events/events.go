package events

import (
	"context"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/interfaces"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"

	uuid "github.com/satori/go.uuid"
)

type EventService struct {
	Processor interfaces.EventStorageProcessor
}

func (es *EventService) GetEvent(ctx context.Context, id uuid.UUID) (models.Event, error) {
	return es.Processor.GetEventByID(ctx, id)
}

func (es *EventService) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	return es.Processor.CreateEvent(ctx, event)
}

func (es *EventService) UpdateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	return es.Processor.UpdateEventByID(ctx, event.EventID, event)
}

func (es *EventService) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return es.Processor.DeleteEventById(ctx, id)
}

func (es *EventService) GetEventsList(ctx context.Context, name string) ([]models.Event, error) {
	return es.Processor.GetUserEvents(ctx, name)
}

func (es *EventService) DeleteUserEvents(ctx context.Context, name string) error {
	return es.Processor.DeleteAllUserEvents(ctx, name)
}
