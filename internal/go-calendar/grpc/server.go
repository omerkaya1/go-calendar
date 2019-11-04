package grpc

import (
	"net"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/services/events"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	gca "github.com/omerkaya1/go-calendar/internal/go-calendar/grpc/go-calendar-api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GoCalendarServer struct {
	Cfg          *config.Config
	Logger       *zap.Logger
	EventService *events.EventService
}

func NewServer(cfg *config.Config, log *zap.Logger, es *events.EventService) *GoCalendarServer {
	return &GoCalendarServer{
		Cfg:          cfg,
		Logger:       log,
		EventService: es,
	}
}

func (s *GoCalendarServer) Run() {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		s.Logger.Sugar().Errorf("%s", err)
	}

	gca.RegisterGoCalendarServerServer(server, s)

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", server.Serve(l))
}
