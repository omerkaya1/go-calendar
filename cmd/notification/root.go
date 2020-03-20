package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/mq"
	"github.com/omerkaya1/go-calendar/internal/notification/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// Handle interrupt signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	defer close(stopChan)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for range stopChan {
			cancel()
			log.Println("context cancellation triggered. The programme's about to stop...")
			return
		}
	}()

	messageQueue, err := mq.NewRabbitMQService(conf.Queue, nil, messageCounter)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
	log.Println("Notification service initialisation")
	if err := messageQueue.EmmitMessages(ctx); err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
	log.Print("Buy!")
}
