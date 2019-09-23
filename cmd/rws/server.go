package rws

import (
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/rws"
	gcl "github.com/omerkaya1/go-calendar/log"
	"github.com/spf13/cobra"
	"log"
)

var (
	cfgPath  string
	host     string
	port     string
	logLevel int
)

var ServerCmd = &cobra.Command{
	Use:   "rws-server",
	Short: "Run RESTful Web Service Server",
	Run: func(cmd *cobra.Command, args []string) {
		// Config-related part
		cfg := &conf.Config{}
		var err error

		if cfgPath != "" {
			cfg, err = cfgFromFile()
		} else {
			cfgFromCmdParams(cfg)
		}

		if err != nil {
			log.Fatalf("InitConfig failed: %v", err)
		}
		// Logger-related part
		logger, err := gcl.InitLogger(cfg.LogLevel)
		if err != nil {
			log.Fatalf("InitLogger failed: %v", err)
		}
		// Init RWS server
		srv := rws.NewServer(cfg, logger)
		srv.Run()
	},
}

func init() {
	// NOTE: I'll have to decide how to properly handle server initialisation
	ServerCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	ServerCmd.PersistentFlags().IntVarP(&logLevel, "log", "l", 1, "changes log level")
	ServerCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "host address")
	ServerCmd.PersistentFlags().StringVarP(&port, "port", "p", "7070", "host port")
}

func cfgFromFile() (*conf.Config, error) {
	cfg, err := conf.InitConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func cfgFromCmdParams(cfg *conf.Config) {
	cfg.LogLevel = logLevel
	cfg.Host = host
	cfg.Port = port
}
