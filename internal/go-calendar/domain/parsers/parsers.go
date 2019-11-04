package parsers

import (
	"time"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/validators"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	gca "github.com/omerkaya1/go-calendar/internal/go-calendar/grpc/go-calendar-api"
	"github.com/satori/go.uuid"
)

// ParseTime is a helper function that converts a time.Time object to a proto Event time object
func ParseTimeToProto(t *time.Time) (*timestamp.Timestamp, error) {
	if t == nil {
		return nil, errors.ErrMalformedTimeObject
	}
	ts, err := ptypes.TimestampProto(*t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// ParseProtoToTime is a helper function that converts a proto Event time object to a time.Time object
func ParseProtoToTime(t *timestamp.Timestamp) (*time.Time, error) {
	if t == nil {
		return nil, errors.ErrMalformedTimeObject
	}
	st, err := ptypes.Timestamp(t)
	if err != nil {
		return nil, err
	}
	return &st, nil
}

// MapToProtoResponseWithID is a helper function that converts an Event internal ID object to a GRPC EventID response
func MapToProtoResponseWithID(id uuid.UUID, err error) *gca.ResponseWithEventID {
	if err != nil {
		return &gca.ResponseWithEventID{
			Result: &gca.ResponseWithEventID_Error{
				Error: err.Error(),
			},
		}
	}
	return &gca.ResponseWithEventID{
		Result: &gca.ResponseWithEventID_EventID{
			EventID: id.String(),
		},
	}
}

// MapToProtoResponseSuccess is a helper function that converts a response object to a GRPC Success response
func MapToProtoResponseSuccess(response string, err error) *gca.ResponseSuccess {
	if err != nil {
		return &gca.ResponseSuccess{
			Result: &gca.ResponseSuccess_Error{
				Error: err.Error(),
			},
		}
	}
	return &gca.ResponseSuccess{
		Result: &gca.ResponseSuccess_Response{
			Response: response,
		},
	}
}

// MapToProtoResponseWithEvent is a helper function that converts an Event object to a GRPC Event response
func MapToProtoResponseWithEvent(event *models.Event, err error) *gca.ResponseWithEvent {
	if err != nil {
		return &gca.ResponseWithEvent{
			Result: &gca.ResponseWithEvent_Error{
				Error: err.Error(),
			},
		}
	}

	startTime, err := ParseTimeToProto(event.StartTime)
	if err != nil {
		return &gca.ResponseWithEvent{
			Result: &gca.ResponseWithEvent_Error{
				Error: err.Error(),
			},
		}
	}

	endTime, err := ParseTimeToProto(event.EndTime)
	if err != nil {
		return &gca.ResponseWithEvent{
			Result: &gca.ResponseWithEvent_Error{
				Error: err.Error(),
			},
		}
	}

	return &gca.ResponseWithEvent{
		Result: &gca.ResponseWithEvent_Event{
			Event: &gca.Event{
				EventId:   event.EventID.String(),
				UserName:  event.UserName,
				EventName: event.EventName,
				Note:      event.Note,
				StartTime: startTime,
				EndTime:   endTime,
			},
		},
	}
}

// TODO: get rid of the code duplication!
// MapProtoEventToEvent is a helper function that converts GRPC CreateEvent request to Event
func MapProtoEventToEvent(cr *gca.CreateEventRequest) (*models.Event, error) {
	startTime, err := ParseProtoToTime(cr.GetStartTime())
	if err != nil {
		return nil, err
	}

	endTime, err := ParseProtoToTime(cr.GetEndTime())
	if err != nil {
		return nil, err
	}
	validators.ValidateTime(startTime, endTime)

	return models.NewEvent(cr.GetUserName(), cr.GetEventName(), cr.GetText(), startTime, endTime), nil
}

// MapProtoEventToOldEvent is a helper function that converts GRPC Event request to Event
func MapProtoEventToOldEvent(er *gca.Event) (*models.Event, error) {
	startTime, err := ParseProtoToTime(er.GetStartTime())
	if err != nil {
		return nil, err
	}

	endTime, err := ParseProtoToTime(er.GetEndTime())
	if err != nil {
		return nil, err
	}
	validators.ValidateTime(startTime, endTime)

	id, err := validators.ValidateID(er.GetEventId())
	if err != nil {
		return nil, err
	}

	return &models.Event{
		EventID:   id,
		UserName:  er.GetUserName(),
		EventName: er.GetEventName(),
		Note:      er.GetNote(),
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}
