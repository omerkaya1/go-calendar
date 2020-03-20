package prometheus

import (
	"errors"
	"fmt"
	grpcp "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Monitor struct {
	Server   *http.Server
	Registry *prometheus.Registry
	Metrics  *grpcp.ServerMetrics
}

func NewMonitor(conf config.PromConf) (*Monitor, error) {
	if !conf.Verify() {
		return nil, errors.New("bad configuration")
	}
	reg := prometheus.NewRegistry()

	server := &http.Server{
		Addr:      fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		TLSConfig: nil,
	}
	http.Handle("/metrics", promhttp.Handler())

	// Create some standard server metrics.
	grpcMetrics := grpcp.NewServerMetrics()
	if err := reg.Register(grpcMetrics); err != nil {
		return nil, err
	}

	return &Monitor{
		Server:   server,
		Registry: reg,
		Metrics:  grpcMetrics,
	}, nil
}
