package rws

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/interfaces"
	"go.uber.org/zap"
)

type Server struct {
	Cfg          *config.Config
	Logger       *zap.Logger
	EventStorage interfaces.EventStorageProcessor
}

func NewServer(cfg *config.Config, log *zap.Logger, es interfaces.EventStorageProcessor) *Server {
	return &Server{
		Cfg:          cfg,
		Logger:       log,
		EventStorage: es,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	r.HandleFunc(path.Join(ApiPrefix, ApiVersion, EventURL), s.CreateEvent).Methods(http.MethodPost)
	r.HandleFunc(path.Join(ApiPrefix, ApiVersion, EventURL, "/{user}", "/{id:[-A-Z0-9a-z]+}"), s.GetEvent).Methods(http.MethodGet)
	r.HandleFunc(path.Join(ApiPrefix, ApiVersion, EventURL, "/{user}"), s.GetUserEvents).Methods(http.MethodGet)
	r.HandleFunc(path.Join(ApiPrefix, ApiVersion, EventURL), s.UpdateEvent).Methods(http.MethodPut)
	r.HandleFunc(path.Join(ApiPrefix, ApiVersion, EventURL, "/{user:[-A-Z0-9a-z]+}", "/{id:[-A-Z0-9a-z]+}"), s.DeleteEvent).Methods(http.MethodDelete)
	r.HandleFunc(path.Join(ApiPrefix, ApiVersion, EventURL, "/{user}"), s.DeleteExpiredEvents).Methods(http.MethodDelete)

	s.Logger.Sugar().Infof("Server initialised on address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", http.ListenAndServe(s.Cfg.Host+":"+s.Cfg.Port, r))
}
