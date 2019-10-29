package grpc

import (
	"context"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/parsers"
	"github.com/omerkaya1/go-calendar/internal/domain/validators"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
	"github.com/satori/go.uuid"
)

// CreateEvent handles GRPC requests to create events
func (s *GoCalendarServer) CreateEvent(ctx context.Context, req *gca.CreateEventRequest) (*gca.ResponseWithEventID, error) {
	// ProtoEvent to Event wrapper
	event, err := parsers.MapProtoEventToEvent(req)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return parsers.MapToProtoResponseWithID(uuid.UUID{}, err), nil
	}
	// Request to the EventService to create an event
	eventID, err := s.EventService.CreateEvent(ctx, event)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			return parsers.MapToProtoResponseWithID(uuid.UUID{}, berr), nil
		}
		return nil, err
	}
	// Log the result and return
	s.Logger.Sugar().Infof("created a new event with id %s", eventID)
	return parsers.MapToProtoResponseWithID(eventID, nil), nil
}

// GetEvent handles GRPC requests to retrieve events
func (s *GoCalendarServer) GetEvent(ctx context.Context, req *gca.RequestEventByID) (*gca.ResponseWithEvent, error) {
	id, err := validators.ValidateID(req.EventID)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}
	event, err := s.EventService.GetEvent(ctx, id)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			return parsers.MapToProtoResponseWithEvent(nil, berr), nil
		}
		return nil, err
	}

	return parsers.MapToProtoResponseWithEvent(&event, nil), nil
}

// UpdateEvent handles GRPC requests to update existing events
func (s *GoCalendarServer) UpdateEvent(ctx context.Context, req *gca.Event) (*gca.ResponseWithEventID, error) {
	// ProtoEvent to Event wrapper
	event, err := parsers.MapProtoEventToOldEvent(req)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return parsers.MapToProtoResponseWithID(uuid.UUID{}, err), nil
	}
	// Request to the EventService to update an event
	eventID, err := s.EventService.UpdateEvent(ctx, event)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			return parsers.MapToProtoResponseWithID(uuid.UUID{}, berr), nil
		}
		return nil, err
	}
	// Log the result and return
	s.Logger.Sugar().Infof("the event with id %s was successfully updated", eventID.String())
	return parsers.MapToProtoResponseWithID(eventID, nil), nil
}

// DeleteEvent handles GRPC requests to delete existing events
func (s *GoCalendarServer) DeleteEvent(ctx context.Context, req *gca.RequestEventByID) (*gca.ResponseSuccess, error) {
	id, err := validators.ValidateID(req.EventID)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}
	if err := s.EventService.DeleteEvent(ctx, id); err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			return parsers.MapToProtoResponseSuccess("", berr), nil
		}
		return nil, err
	}
	s.Logger.Sugar().Infof("the event with id %s was successfully deleted", id)
	return parsers.MapToProtoResponseSuccess("event was successfully deleted", nil), nil
}