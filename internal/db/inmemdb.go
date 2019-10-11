package db

import (
	"context"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/satori/go.uuid"
	"sync"
)

// InMemoryEventStorage represents an object that holds all created events in memory.
type InMemoryEventStorage struct {
	m  *sync.RWMutex
	db map[uuid.UUID]models.Event
}

// NewInMemoryEventStorage returns a new InMemoryEventStorage object
func NewInMemoryEventStorage() (*InMemoryEventStorage, error) {
	return &InMemoryEventStorage{db: make(map[uuid.UUID]models.Event), m: &sync.RWMutex{}}, nil
}

// TODO: add comments!
// GetEventByID .
func (imes *InMemoryEventStorage) GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	imes.m.RLock()
	defer imes.m.RUnlock()
	if event, ok := imes.db[id]; !ok {
		return models.Event{}, errors.ErrEventDoesNotExist
	} else {
		return event, nil
	}
}

// CreateEvent .
func (imes *InMemoryEventStorage) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	imes.m.Lock()
	defer imes.m.Unlock()
	if err := imes.checkEventCollision(event); err != nil {
		return uuid.UUID{}, err
	}
	imes.db[event.EventId] = *event
	return event.EventId, nil
}

// DeleteEventById .
func (imes *InMemoryEventStorage) DeleteEventById(ctx context.Context, id uuid.UUID) error {
	imes.m.Lock()
	defer imes.m.Unlock()
	if event, ok := imes.db[id]; !ok {
		return errors.ErrEventDoesNotExist
	} else {
		delete(imes.db, event.EventId)
	}
	return nil
}

// UpdateEventByID .
func (imes *InMemoryEventStorage) UpdateEventByID(ctx context.Context, id uuid.UUID, event *models.Event) (uuid.UUID, error) {
	imes.m.Lock()
	defer imes.m.Unlock()
	if oldEvent, ok := imes.db[id]; !ok {
		return id, errors.ErrEventDoesNotExist
	} else {
		updated := models.ComposeEvent(oldEvent, event)
		if err := imes.checkEventCollision(updated); err != nil {
			return id, err
		}
		delete(imes.db, oldEvent.EventId)
		imes.db[event.EventId] = *updated
		return updated.EventId, nil
	}
}

// UpdateEventByName .
func (imes *InMemoryEventStorage) UpdateEventByName(ctx context.Context, eventName string, event *models.Event) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

// GetUserEvents .
func (imes *InMemoryEventStorage) GetUserEvents(ctx context.Context, user string) ([]models.Event, error) {
	return []models.Event{}, nil
}

// GetEventByName .
func (imes *InMemoryEventStorage) GetEventByName(ctx context.Context, name string) (models.Event, error) {
	return models.Event{}, nil
}

// DeleteAllUserEvents .
func (imes *InMemoryEventStorage) DeleteAllUserEvents(ctx context.Context, user string) error {
	return nil
}

func (imes *InMemoryEventStorage) checkEventCollision(event *models.Event) error {
	for _, v := range imes.db {
		// A new event takes place within the time interval of another event
		if v.StartTime.Before(*event.StartTime) && v.EndTime.After(*event.EndTime) {
			return errors.ErrEventCollisionInInterval
		}
		// A new event takes place within the time interval of another event
		if v.StartTime.After(*event.StartTime) && v.EndTime.Before(*event.EndTime) {
			return errors.ErrEventCollisionOutInterval
		}
		// A new event takes place within the time interval of another event
		if v.StartTime.Equal(*event.StartTime) || v.EndTime.Equal(*event.EndTime) {
			return errors.ErrEventCollisionMatch
		}
	}
	return nil
}
