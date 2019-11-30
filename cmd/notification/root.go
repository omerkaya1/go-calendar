package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/mq"
	"github.com/omerkaya1/go-calendar/internal/notification/errors"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:     "notification",
	Short:   "simple notification service that queries RabbitMQ for messages that belong to a particular message queue",
	Example: "# Initialise from configuration file\n notification -c /path/to/config.json",
	Run:     startNotificationService,
}

var (
	// Create a customized counter metric.
	messageCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "rabbit_mq_sent_messages",
		Help: "Total number of messages sent by the MQ service",
	})
)

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "specifies the path to a configuration file")
	prometheus.MustRegister(messageCounter)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("MessageQueue metrics initialisation...")
		log.Fatal(http.ListenAndServe("notification:9187", nil))
	}()
}

func startNotificationService(cmd *cobra.Command, args []string) {
	if configFile == "" {
		log.Fatalf("%s: %s", errors.ErrCMDPrefix, errors.ErrBadConfigFile)
	}

	conf, err := config.InitConfig(configFile)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrCMDPrefix, err)
	}

	messageQueue, err := mq.NewRabbitMQService(conf, nil, messageCounter)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
	log.Println("Notification service initialisation")
	if err := messageQueue.EmmitMessages(); err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
}
