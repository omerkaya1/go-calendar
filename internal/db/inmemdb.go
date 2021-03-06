package db

import (
	"context"
	"log"
	"sync"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"

	"github.com/satori/go.uuid"
)

// InMemoryEventStorage represents an object that holds all created events in memory.
type InMemoryEventStorage struct {
	m  *sync.RWMutex
	db map[uuid.UUID]models.Event
}

// NewInMemoryEventStorage returns a new InMemoryEventStorage object.
func NewInMemoryEventStorage() (*InMemoryEventStorage, error) {
	return &InMemoryEventStorage{db: make(map[uuid.UUID]models.Event), m: &sync.RWMutex{}}, nil
}

// GetEventByID returns a requested by its ID event or an error on failure
func (imes *InMemoryEventStorage) GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	imes.m.RLock()
	defer imes.m.RUnlock()
	if event, ok := imes.db[id]; !ok {
		return models.Event{}, errors.ErrEventDoesNotExist
	} else {
		return event, nil
	}
}

// CreateEvent creates new event and stores it to the DB and returns the event's internal ID
// On failure it returns an empty ID object and an error
func (imes *InMemoryEventStorage) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	imes.m.Lock()
	defer imes.m.Unlock()
	if err := imes.checkEventCollision(event); err != nil {
		return uuid.UUID{}, err
	}
	imes.db[event.EventID] = *event
	return event.EventID, nil
}

// DeleteEventById deletes an existing event, which ID was passed as an argument.
func (imes *InMemoryEventStorage) DeleteEventById(ctx context.Context, id uuid.UUID) error {
	imes.m.Lock()
	defer imes.m.Unlock()
	if event, ok := imes.db[id]; !ok {
		return errors.ErrEventDoesNotExist
	} else {
		delete(imes.db, event.EventID)
	}
	return nil
}

// UpdateEventByID updates an existing event, which ID was passed as an argument
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
		delete(imes.db, oldEvent.EventID)
		imes.db[event.EventID] = *updated
		return updated.EventID, nil
	}
}

// GetUpcomingEvents .
func (imes *InMemoryEventStorage) GetUpcomingEvents(ctx context.Context) ([]models.Event, error) {
	log.Println("Implement me!")
	return nil, nil
}

// GetUserEvents .
func (imes *InMemoryEventStorage) GetUserEvents(ctx context.Context, user string) ([]models.Event, error) {
	log.Println("Implement me!")
	return []models.Event{}, nil
}

// DeleteAllUserEvents .
func (imes *InMemoryEventStorage) DeleteExpiredEvents(ctx context.Context, user string) (int64, error) {
	log.Println("Implement me!")
	return 0, nil
}

func (imes *InMemoryEventStorage) checkEventCollision(event *models.Event) error {
	if len(imes.db) == 0 {
		return nil
	}
	for _, v := range imes.db {
		// A new event takes place within the time interval of another event
		if event.EndTime.Before(*v.StartTime) || event.StartTime.After(*v.EndTime) {
			return nil
		}
	}
	return errors.ErrEventCollisionInInterval
}
