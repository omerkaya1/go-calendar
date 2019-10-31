package rws

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/omerkaya1/go-calendar/internal/domain/validators"
)

const (
	RWSApiPrefix  = "/api"
	RWSapiVersion = "/v1"
	RWSeventURL   = "/event"
)

// CreateEvent handles RWS requests to create events
func (s *Server) CreateEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	event := &models.EventJSON{}

	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	startTime, err := validators.ValidateDate(event.StartTime)
	if err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	endTime, err := validators.ValidateDate(event.EndTime)
	if err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	req := models.NewEvent(event.UserName, event.EventName, event.Note, startTime, endTime)

	eventID, err := s.EventService.CreateEvent(context.Background(), req)
	if err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	s.Logger.Sugar().Infof("created a new event with id %s", eventID)
	rw.Write([]byte(eventID.String()))
	return
}

// UpdateEvent handles RWS requests to update existing events
func (s *Server) UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	event := &models.EventJSON{}
	var err error

	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := validators.ValidateID(event.EventID)
	if err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	start, finish := &time.Time{}, &time.Time{}
	if event.StartTime != "" && event.EndTime != "" {
		start, err = validators.ValidateDate(event.StartTime)
		if err != nil {
			s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
			return
		}
		finish, err = validators.ValidateDate(event.EndTime)
		if err != nil {
			s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
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
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	s.Logger.Sugar().Infof("updated the event with id %s", eventID)
	rw.Write([]byte(eventID.String()))
	return
}

// DeleteEvent handles RWS requests to delete existing events
func (s *Server) DeleteEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	id, err := validators.ValidateID(params["id"])
	if err != nil {
		s.formResponse(rw, true, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.EventService.DeleteEvent(context.Background(), id); err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	s.Logger.Sugar().Infof("the event with id %s was successfully deleted", id.String())
	rw.Write([]byte("event was successfully deleted"))
	return
}

// GetEvent handles RWS requests to retrieve events
func (s *Server) GetEvent(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	id, err := validators.ValidateID(params["id"])
	if err != nil {
		s.formResponse(rw, true, http.StatusBadRequest, err.Error())
		return
	}

	event, err := s.EventService.GetEvent(context.Background(), id)
	if err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := json.Marshal(event)
	if err != nil {
		s.formResponse(rw, true, http.StatusInternalServerError, err.Error())
		return
	}

	rw.Write(resp)
	return
}

func (s *Server) formResponse(rw http.ResponseWriter, error bool, code int, message string) {
	if error {
		s.Logger.Sugar().Error(message)
		rw.WriteHeader(code)
	}
	rw.Write([]byte(message))
}
