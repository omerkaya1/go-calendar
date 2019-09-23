package rws

import (
	"github.com/gorilla/mux"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	Cfg    *conf.Config
	Logger *zap.Logger
	// TODO: replace with a real DB implementation
	dummDB []models.Event
}

func NewServer(cfg *conf.Config, log *zap.Logger) *Server {
	return &Server{
		Cfg:    cfg,
		Logger: log,
		dummDB: make([]models.Event, 100, 100),
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	r.HandleFunc(apiPrefix+apiVersion+getEventURL, s.getEvent)
	r.HandleFunc(apiPrefix+apiVersion+createEventURL, s.createEvent)
	r.HandleFunc(apiPrefix+apiVersion+updateEventURL, s.updateEvent)
	r.HandleFunc(apiPrefix+apiVersion+deleteEventURL, s.deleteEvent)

	//s.Logger.Sugar().Info("Server initialisation...\n")
	s.Logger.Sugar().Infof("Server initialised on address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%v", http.ListenAndServe(s.Cfg.Host+":"+s.Cfg.Port, r))
}
