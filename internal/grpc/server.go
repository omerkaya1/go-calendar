package grpc

import (
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/services"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
	queue "github.com/omerkaya1/go-calendar/internal/message-queue"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GoCalendarServer struct {
	Cfg          *conf.Config
	Logger       *zap.Logger
	EventService *services.EventService
	Queue        *queue.MessageQueue
}

func NewServer(cfg *conf.Config, log *zap.Logger, es *services.EventService, q *queue.MessageQueue) *GoCalendarServer {
	return &GoCalendarServer{
		Cfg:          cfg,
		Logger:       log,
		EventService: es,
		Queue:        q,
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
