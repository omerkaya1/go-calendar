package rws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/validators"
)

const (
	ApiPrefix  = "/api"
	ApiVersion = "/v1"
	EventURL   = "/event"
)

// GetEvent handles RWS requests to retrieve events
func (s *Server) GetEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if idParam, ok := params["id"]; ok {
		id, err := validators.ValidateID(idParam)
		if err != nil {
			s.formResponse(rw, err, http.StatusBadRequest, "")
			return
		}

		event, err := s.EventService.GetEvent(context.Background(), id)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}

		resp, err := json.Marshal(event)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}
		if n, err := rw.Write(resp); err != nil {
			s.Logger.Sugar().Infof("%s: %s, %d bytes were written", errors.ErrAPIPrefix, err, n)
		}
		return
	}
	s.formResponse(rw, errors.ErrURLPathParameters, http.StatusBadRequest, "")
}

// CreateEvent handles RWS requests to create events
func (s *Server) CreateEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	event := &models.EventJSON{}

	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}

	startTime, err := validators.ValidateDate(event.StartTime)
	if err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}

	endTime, err := validators.ValidateDate(event.EndTime)
	if err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}

	req := models.NewEvent(event.UserName, event.EventName, event.Note, startTime, endTime)

	eventID, err := s.EventService.CreateEvent(context.Background(), req)
	if err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}
	// Log result and send response
	s.formResponse(rw, nil, http.StatusCreated, fmt.Sprintf("created a new event with id %s", eventID))
}

// UpdateEvent handles RWS requests to update existing events
func (s *Server) UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	event := &models.EventJSON{}
	var err error

	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}

	id, err := validators.ValidateID(event.EventID)
	if err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}

	start, finish := &time.Time{}, &time.Time{}
	if event.StartTime != "" && event.EndTime != "" {
		start, err = validators.ValidateDate(event.StartTime)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}
		finish, err = validators.ValidateDate(event.EndTime)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}
	}

	req := &models.Event{
		EventID:   id,
		UserName:  event.UserName,
		EventName: event.EventName,
		Note:      event.Note,
		StartTime: start,
		EndTime:   finish,
	}

	eventID, err := s.EventService.UpdateEvent(context.Background(), req)
	if err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}
	s.formResponse(rw, nil, http.StatusOK, fmt.Sprintf("updated the event with id %s", eventID))
}

// DeleteEvent handles RWS requests to delete existing events
func (s *Server) DeleteEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	id, err := validators.ValidateID(params["id"])
	if err != nil {
		s.formResponse(rw, err, http.StatusBadRequest, "")
		return
	}

	if err := s.EventService.DeleteEvent(context.Background(), id); err != nil {
		s.formResponse(rw, err, http.StatusInternalServerError, "")
		return
	}
	s.formResponse(rw, err, http.StatusInternalServerError, fmt.Sprintf("the event with id %s was deleted", id))
}

// DeleteEvent handles RWS requests to delete existing events
func (s *Server) DeleteExpiredEvents(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if user, ok := params["user"]; ok {
		n, err := s.EventService.DeleteOldEvents(context.Background(), user)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}
		s.formResponse(rw, nil, http.StatusOK, fmt.Sprintf("%d old events belonging to %s were deleted", n, user))
		return
	}
	s.formResponse(rw, errors.ErrURLPathParameters, http.StatusBadRequest, "")
}

// GetEvent handles RWS requests to retrieve events
func (s *Server) GetUserEvents(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if user, ok := params["user"]; ok {
		event, err := s.EventService.GetEventsList(context.Background(), user)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}

		resp, err := json.Marshal(event)
		if err != nil {
			s.formResponse(rw, err, http.StatusInternalServerError, "")
			return
		}
		if n, err := rw.Write(resp); err != nil {
			s.Logger.Sugar().Infof("%s: %s, %d bytes were written", errors.ErrAPIPrefix, err, n)
		}
		return
	}
	s.formResponse(rw, errors.ErrURLPathParameters, http.StatusBadRequest, "")
}

func (s *Server) formResponse(rw http.ResponseWriter, err error, code int, message string) {
	if err != nil {
		s.Logger.Sugar().Error(err)
		rw.WriteHeader(code)
	}
	s.Logger.Sugar().Info(message)
	if n, err := rw.Write([]byte(message)); err != nil {
		s.Logger.Sugar().Infof("%s: %s, %d bytes were written", errors.ErrAPIPrefix, err, n)
	}
}
