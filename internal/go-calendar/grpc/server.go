package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/services/events"
	gca "github.com/omerkaya1/go-calendar/internal/go-calendar/grpc/api"
	"github.com/omerkaya1/go-calendar/internal/prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GoCalendarServer struct {
	Cfg          *config.Config
	Logger       *zap.Logger
	EventService *events.EventService
	Monitoring   *prometheus.Monitor
}

func NewServer(cfg *config.Config, log *zap.Logger, es *events.EventService, m *prometheus.Monitor) *GoCalendarServer {
	return &GoCalendarServer{
		Cfg:          cfg,
		Logger:       log,
		EventService: es,
		Monitoring:   m,
	}
}

func (s *GoCalendarServer) Run() {
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		s.Logger.Sugar().Errorf("%s", err)
	}

	gca.RegisterGoCalendarServerServer(server, s)

	s.Monitoring.Metrics.EnableHandlingTimeHistogram()
	s.Monitoring.Metrics.InitializeMetrics(server)

	// Start your http server for prometheus.
	go func() {
		if err := s.Monitoring.Server.ListenAndServe(); err != nil {
			s.Logger.Sugar().Fatal("Unable to start a http server.")
		}
	}()

	// Listen for the OS signals
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	// Gracefully shut down the server should the proper os signal arrive
	go func() {
		for range exit {
			s.Logger.Sugar().Info("Interrupt signal received.")
			server.GracefulStop()
			s.Logger.Sugar().Info("Graceful shutdown performed. Bye!")
			return
		}
	}()

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", server.Serve(l))
}
