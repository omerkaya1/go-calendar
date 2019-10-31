package services

import (
	"context"

	"github.com/omerkaya1/go-calendar/internal/domain/interfaces"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
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
	if event.EventID.String() != "" {
		return es.Processor.UpdateEventByID(ctx, event.EventID, event)
	}
	return es.Processor.UpdateEventByName(ctx, event.EventName, event)
}

func (es *EventService) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return es.Processor.DeleteEventById(ctx, id)
}

func (es *EventService) GetEventsList(ctx context.Context) {

}
