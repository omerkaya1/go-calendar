package grpc

import (
	"net"

	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/services"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GoCalendarServer struct {
	Cfg          *conf.Config
	Logger       *zap.Logger
	EventService *services.EventService
}

func NewServer(cfg *conf.Config, log *zap.Logger, es *services.EventService) *GoCalendarServer {
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
