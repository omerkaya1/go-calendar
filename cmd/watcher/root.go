package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/omerkaya1/go-calendar/internal/db"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/mq"
	"github.com/omerkaya1/go-calendar/internal/watcher/errors"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:     "watcher",
	Short:   "simple DB watcher service that queries a DB with calendar events and enqueues them in a message queue",
	Example: "  watcher -c /path/to/config.json",
	Run:     startNotificationService,
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "specifies the path to a configuration file")
}

func startNotificationService(cmd *cobra.Command, args []string) {
	if configFile == "" {
		log.Fatalf("%s:%s", errors.ErrCMDPrefix, errors.ErrBadConfigFile)
	}

	conf, err := config.InitConfig(configFile)
	if err != nil {
		log.Fatalf("%s:%s", errors.ErrCMDPrefix, err)
	}

	esp, err := db.NewMainEventStorage(conf.DB)
	if err != nil {
		log.Fatalf("%s:%s", errors.ErrDBPrefix, err)
	}

	messageQueue, err := mq.NewRabbitMQService(conf.Queue, esp, nil)
	if err != nil {
		log.Fatalf("%s:%s", errors.ErrMQPrefix, err)
	}

	// Handle interrupt signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	defer close(stopChan)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for range stopChan {
			cancel()
			log.Print("context cancellation triggered. The programme's about to stop...")
			return
		}
	}()

	log.Println("Watcher service initialisation")
	if err := messageQueue.ProduceMessages(ctx); err != nil {
		log.Fatalf("%s:%s", errors.ErrMQPrefix, err)
	}
	log.Print("Buy!")
}
