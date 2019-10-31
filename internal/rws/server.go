package rws

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/services"
	"go.uber.org/zap"
)

type Server struct {
	Cfg          *conf.Config
	Logger       *zap.Logger
	EventService *services.EventService
}

func NewServer(cfg *conf.Config, log *zap.Logger, es *services.EventService) *Server {
	return &Server{
		Cfg:          cfg,
		Logger:       log,
		EventService: es,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL), s.CreateEvent).Methods(http.MethodPost)
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL, "/{id:[-A-Z0-9a-z]+}"), s.GetEvent).Methods(http.MethodGet)
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL), s.UpdateEvent).Methods(http.MethodPut)
	r.HandleFunc(path.Join(RWSApiPrefix, RWSapiVersion, RWSeventURL, "/{id:[-A-Z0-9a-z]+}"), s.DeleteEvent).Methods(http.MethodDelete)

	s.Logger.Sugar().Infof("Server initialised on address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", http.ListenAndServe(s.Cfg.Host+":"+s.Cfg.Port, r))
}
