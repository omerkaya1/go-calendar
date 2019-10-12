package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/omerkaya1/go-calendar/internal/domain/parsers"
	"github.com/omerkaya1/go-calendar/internal/domain/validators"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
	"time"
)

func (s *GoCalendarServer) CreateEvent(ctx context.Context, req *gca.CreateEventRequest) (*gca.ResponseWithEventID, error) {
	// TODO: Remains for later consideration.
	//userName := ""
	//if o := ctx.Value("user_name"); o != nil {
	//	userName, _ = o.(string)
	//}

	startTime := (*time.Time)(nil)
	if req.GetStartTime() != nil {
		st, err := ptypes.Timestamp(req.GetStartTime())
		if err != nil {
			return nil, err
		}
		startTime = &st
	}

	endTime := (*time.Time)(nil)
	if req.GetEndTime() != nil {
		et, err := ptypes.Timestamp(req.GetEndTime())
		if err != nil {
			return nil, err
		}
		endTime = &et
	}

	eventID, err := s.EventService.CreateEvent(ctx, models.NewEvent(req.GetUserName(), req.GetEventName(), req.GetText(), startTime, endTime))
	if err != nil {
		if berr, ok := err.(errors.GoCalendarError); ok {
			resp := &gca.ResponseWithEventID{
				Result: &gca.ResponseWithEventID_Error{
					Error: string(berr),
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
	id := validators.ValidateID(req.EventID)
	event, err := s.EventService.GetEvent(ctx, id)
	if err != nil {
		if berr, ok := err.(errors.GoCalendarError); ok {
			resp := &gca.ResponseWithEvent{
				Result: &gca.ResponseWithEvent_Error{
					Error: string(berr),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}

	resp := &gca.ResponseWithEvent{
		Result: &gca.ResponseWithEvent_Event{
			Event: &gca.Event{
				EventId:   event.EventId.String(),
				UserName:  event.UserName,
				EventName: event.EventName,
				Note:      event.Note,
				StartTime: parsers.ParseTime(event.StartTime),
				EndTime:   parsers.ParseTime(event.EndTime),
			},
		},
	}
	return resp, nil
}

func (s *GoCalendarServer) UpdateEvent(ctx context.Context, req *gca.Event) (*gca.ResponseWithEventID, error) {
	startTime := (*time.Time)(nil)
	if req.GetStartTime() != nil {
		st, err := ptypes.Timestamp(req.GetStartTime())
		if err != nil {
			return nil, err
		}
		startTime = &st
	}

	endTime := (*time.Time)(nil)
	if req.GetEndTime() != nil {
		et, err := ptypes.Timestamp(req.GetEndTime())
		if err != nil {
			return nil, err
		}
		endTime = &et
	}

	newEvent := &models.Event{
		EventId:   validators.ValidateID(req.GetEventId()),
		UserName:  req.GetUserName(),
		EventName: req.GetEventName(),
		Note:      req.GetNote(),
		StartTime: startTime,
		EndTime:   endTime}

	eventID, err := s.EventService.UpdateEvent(ctx, newEvent)
	if err != nil {
		if berr, ok := err.(errors.GoCalendarError); ok {
			resp := &gca.ResponseWithEventID{
				Result: &gca.ResponseWithEventID_Error{
					Error: string(berr),
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
	id := validators.ValidateID(req.EventID)
	if err := s.EventService.DeleteEvent(ctx, id); err != nil {
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
			Response: "event with was successfully deleted",
		},
	}
	s.Logger.Sugar().Infof("the event with id %s was successfully deleted", id)
	return resp, nil
}
