package db

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/satori/go.uuid"
	"sync"
)

// MainEventStorage object holds everything related to the DB interactions
type MainEventStorage struct {
	m  *sync.RWMutex
	db *sqlx.DB
}

// NewMainEventStorage returns new MainEventStorage object to the callee
func NewMainEventStorage(cfg conf.DBConf) (*MainEventStorage, error) {
	if cfg.Name == "" || cfg.User == "" || cfg.SSLMode == "" || cfg.Password == "" {
		return nil, errors.ErrBadDBConfiguration
	}
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Password, cfg.User, cfg.Name, cfg.SSLMode)
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return &MainEventStorage{db: db, m: &sync.RWMutex{}}, nil
}

// GetEventByID returns an event requested by the callee
func (edb *MainEventStorage) GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	return edb.getEventByID(ctx, id)
}

// UpdateEventByID updates an event stored in the DB
func (edb *MainEventStorage) UpdateEventByID(ctx context.Context, id uuid.UUID, event *models.Event) (uuid.UUID, error) {
	// Check whether a requested event exists or not
	e, err := edb.getEventByID(ctx, id)
	if err != nil {
		return uuid.UUID{}, err
	}
	newEvent := models.ComposeEvent(e, event)
	// Compose a query
	query := `update events 
			  set title = $1, note = $2, start_time = $3, end_time = $4
			  where id = $5`
	_, err = edb.db.ExecContext(ctx, query, newEvent.EventName, newEvent.Note, newEvent.StartTime, newEvent.EndTime, id.String())
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

// CreateEvent creates a new event and stores it in the DB
func (edb *MainEventStorage) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	// Check whether an event does not take place in the middle of some other event
	if err := edb.checkIntersection(ctx, event); err != nil {
		return uuid.UUID{}, err
	}
	// Compose a DB query
	query := `
		insert into events(id, user_name, title, note, start_time, end_time)
		values (:id, :user_name, :title, :note, :start_time, :end_time)
	`
	_, err := edb.db.NamedExecContext(ctx, query, models.Event{
		EventID:   event.EventID,
		UserName:  event.UserName,
		EventName: event.EventName,
		Note:      event.Note,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	return event.EventID, nil
}

// DeleteEventById deletes an event with the specified ID
func (edb *MainEventStorage) DeleteEventById(ctx context.Context, id uuid.UUID) error {
	query := `delete from events where id=$1`
	result, err := edb.db.ExecContext(ctx, query, id.String())
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.ErrNoOpDBAction
	}
	return nil
}

// GetUpcomingEvents method queries the DB and returns events for the current day and the day after
func (edb *MainEventStorage) GetUpcomingEvents(ctx context.Context) ([]models.Event, error) {
	eventList := make([]models.Event, 0)
	e := models.Event{}
	query := "select * from events where start_time::date=current_date or start_time::date=current_date+interval '1 day'"
	rows, err := edb.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if err := rows.StructScan(&e); err != nil {
				return eventList, err
			}
			eventList = append(eventList, e)
		}
	}
	return eventList, nil
}

// MEMO: for later implementation.
// GetUserEvents .
func (edb *MainEventStorage) GetUserEvents(ctx context.Context, user string) ([]models.Event, error) {
	return nil, nil
}

func (edb *MainEventStorage) GetEventByName(ctx context.Context, name string) (models.Event, error) {
	return models.Event{}, nil
}

// UpdateEventByName .
func (edb *MainEventStorage) UpdateEventByName(ctx context.Context, eventName string, event *models.Event) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (edb *MainEventStorage) DeleteAllUserEvents(ctx context.Context, user string) error {
	return nil
}

func (edb *MainEventStorage) checkIntersection(ctx context.Context, event *models.Event) error {
	// Prepare a query
	query := "select * from events where user_name=$1 order by start_time asc"
	rows, err := edb.db.QueryxContext(ctx, query, event.UserName)
	if err != nil {
		return err
	}
	defer rows.Close()
	// MEMO: is there a better way to check whether a query returned 0 rows?
	exist := false
	e := models.Event{}
	// Iterate over the query results
	for rows.Next() {
		exist = true
		if err := rows.StructScan(&e); err != nil {
			return err
		}
		// If this condition is satisfied, then a new event has no overlaps with other events
		if event.EndTime.Before(*e.StartTime) || event.StartTime.After(*e.EndTime) {
			return nil
		}
	}
	// Check the DB error
	if err := rows.Err(); err != nil {
		return err
	}
	// No events in the DB, exit without failure
	if !exist {
		return nil
	}
	return errors.ErrEventCollisionInInterval
}

func (edb *MainEventStorage) getEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	e := models.Event{}
	query := "select * from events where id = $1"
	return e, edb.db.GetContext(ctx, &e, query, id.String())
}
