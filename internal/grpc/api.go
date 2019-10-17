package grpc

import (
	"context"
	//"github.com/golang/protobuf/ptypes"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/omerkaya1/go-calendar/internal/domain/parsers"
	"github.com/omerkaya1/go-calendar/internal/domain/validators"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
)

func (s *GoCalendarServer) CreateEvent(ctx context.Context, req *gca.CreateEventRequest) (*gca.ResponseWithEventID, error) {

	startTime, err := parsers.ParseProtoToTime(req.GetStartTime())
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	endTime, err := parsers.ParseProtoToTime(req.GetEndTime())
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	eventID, err := s.EventService.CreateEvent(ctx, models.NewEvent(req.GetUserName(), req.GetEventName(), req.GetText(), startTime, endTime))
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			resp := &gca.ResponseWithEventID{
				Result: &gca.ResponseWithEventID_Error{
					Error: berr.Error(),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}

	resp := &gca.ResponseWithEventID{
		Result: &gca.ResponseWithEventID_EventID{
			EventID: eventID.String(),
		},
	}
	s.Logger.Sugar().Infof("created a new event with id %s", resp.GetEventID())
	return resp, nil
}

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
			resp := &gca.ResponseWithEvent{
				Result: &gca.ResponseWithEvent_Error{
					Error: berr.Error(),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}

	startTime, err := parsers.ParseTimeToProto(event.StartTime)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	endTime, err := parsers.ParseTimeToProto(event.EndTime)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	resp := &gca.ResponseWithEvent{
		Result: &gca.ResponseWithEvent_Event{
			Event: &gca.Event{
				EventId:   event.EventId.String(),
				UserName:  event.UserName,
				EventName: event.EventName,
				Note:      event.Note,
				StartTime: startTime,
				EndTime:   endTime,
			},
		},
	}
	return resp, nil
}

func (s *GoCalendarServer) UpdateEvent(ctx context.Context, req *gca.Event) (*gca.ResponseWithEventID, error) {

	startTime, err := parsers.ParseProtoToTime(req.GetStartTime())
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	endTime, err := parsers.ParseProtoToTime(req.GetEndTime())
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	id, err := validators.ValidateID(req.GetEventId())
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}

	newEvent := &models.Event{
		EventId:   id,
		UserName:  req.GetUserName(),
		EventName: req.GetEventName(),
		Note:      req.GetNote(),
		StartTime: startTime,
		EndTime:   endTime}

	eventID, err := s.EventService.UpdateEvent(ctx, newEvent)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			resp := &gca.ResponseWithEventID{
				Result: &gca.ResponseWithEventID_Error{
					Error: berr.Error(),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}

	resp := &gca.ResponseWithEventID{
		Result: &gca.ResponseWithEventID_EventID{
			EventID: eventID.String(),
		},
	}
	s.Logger.Sugar().Infof("the event with id %s was successfully updated", resp.GetEventID())
	return resp, nil
}

func (s *GoCalendarServer) DeleteEvent(ctx context.Context, req *gca.RequestEventByID) (*gca.ResponseSuccess, error) {
	id, err := validators.ValidateID(req.EventID)
	if err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		return nil, err
	}
	if err := s.EventService.DeleteEvent(ctx, id); err != nil {
		s.Logger.Sugar().Errorf("%s: %s", errors.ErrAPIPrefix, err)
		if berr, ok := err.(errors.GoCalendarError); ok {
			resp := &gca.ResponseSuccess{
				Result: &gca.ResponseSuccess_Error{
					Error: string(berr),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}

	resp := &gca.ResponseSuccess{
		Result: &gca.ResponseSuccess_Response{
			Response: "event was successfully deleted",
		},
	}
	s.Logger.Sugar().Infof("the event with id %s was successfully deleted", id)
	return resp, nil
}
