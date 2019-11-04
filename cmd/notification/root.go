package main

import (
	"log"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/mq"
	"github.com/omerkaya1/go-calendar/internal/notification/errors"
	"github.com/spf13/cobra"
)

var configFile string

// RootCmd is the main entry point to the programme
var rootCmd = &cobra.Command{
	Use:     "notification",
	Short:   "simple notification service that queries RabbitMQ for messages that belong to a particular message queue",
	Example: "# Initialise from configuration file notification -c /path/to/config.json",
	Run:     startNotificationService,
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "specifies the path to a configuration file")
}

func startNotificationService(cmd *cobra.Command, args []string) {
	if configFile == "" {
		log.Fatalf("%s: %s", errors.ErrCMDPrefix, errors.ErrBadConfigFile)
	}

	conf, err := config.InitConfig(configFile)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrCMDPrefix, err)
	}

	messageQueue, err := mq.NewRabbitMQService(conf, nil)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
	log.Println("Notification service initialisation")
	if err := messageQueue.EmmitMessages(); err != nil {
		log.Fatalf("%s: %s", errors.ErrMQPrefix, err)
	}
}
