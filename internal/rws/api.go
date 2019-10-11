package rws

import (
	"net/http"
)

const (
	apiPrefix  = "/api"
	apiVersion = "/v1"
	eventURL   = "/event"
)

func (s *Server) CreateEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Info("createEvent triggered")
}

func (s *Server) UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Info("updateEvent triggered")
}

func (s *Server) DeleteEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Info("deleteEvent triggered")
}

func (s *Server) GetEvent(rw http.ResponseWriter, r *http.Request) {
	// TODO: Implement me!
	s.Logger.Info("getEvent triggered")
}
