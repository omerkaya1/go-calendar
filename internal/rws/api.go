package rws

import (
	"net/http"
)

const (
	apiPrefix      = "/"
	apiVersion     = "v1"
	createEventURL = "/create"
	deleteEventURL = "/delete"
	updateEventURL = "/update"
	getEventURL    = "/event"
)

func (s *Server) createEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Warn("createEvent triggered")
}

func (s *Server) updateEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Warn("updateEvent triggered")
}

func (s *Server) deleteEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Warn("deleteEvent triggered")
}

func (s *Server) getEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Warn("getEvent triggered")
}
