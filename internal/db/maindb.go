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
)

type MainEventStorage struct {
	db *sqlx.DB
}

func NewMainEventStorage(cfg conf.DBConf) (*MainEventStorage, error) {
	if cfg.Name == "" || cfg.User == "" || cfg.SSLMode == "" {
		return nil, errors.ErrBadDBConfiguration
	}
	// NOTE: password?
	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s", cfg.User, cfg.Name, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &MainEventStorage{db: db}, nil
}

func (edb *MainEventStorage) GetUserEvents(ctx context.Context, user string) ([]models.Event, error) {
	return []models.Event{}, nil
}

func (edb *MainEventStorage) GetEventByID(ctx context.Context, id uuid.UUID) (models.Event, error) {
	return models.Event{}, nil
}

func (edb *MainEventStorage) GetEventByName(ctx context.Context, name string) (models.Event, error) {
	return models.Event{}, nil
}

func (edb *MainEventStorage) UpdateEventByID(ctx context.Context, id uuid.UUID, event *models.Event) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

// UpdateEventByName .
func (edb *MainEventStorage) UpdateEventByName(ctx context.Context, eventName string, event *models.Event) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (edb *MainEventStorage) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (edb *MainEventStorage) DeleteEventById(ctx context.Context, uuid uuid.UUID) error {
	return nil
}

func (edb *MainEventStorage) DeleteAllUserEvents(ctx context.Context, user string) error {
	return nil
}
