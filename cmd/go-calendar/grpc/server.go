package grpc

import (
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/interfaces"
	"github.com/omerkaya1/go-calendar/internal/prometheus"
	"log"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/services/events"

	"github.com/omerkaya1/go-calendar/internal/db"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/grpc"
	gcl "github.com/omerkaya1/go-calendar/log"
	"github.com/spf13/cobra"
)

var (
	// Server related
	cfgPath, connHost, connPort string
	logLevel                    int
	// DB related
	dbUser, dbName, dbPassword, sslMode, dbHost, dbPort string
)

var ServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "Run GRPC Server",
	Example: `# Initialise from configuration file
go-calendar grpc-server -c /path/to/config.json

# Initialise from parameters
go-calendar grpc-server --host=127.0.0.1 --port=7777 --log=2 --dbname=db_name --dbuser=username`,
	Run: serverStartCmdFunc,
}

func init() {
	ServerCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	ServerCmd.Flags().IntVarP(&logLevel, "log", "l", 1, "changes log level")
	ServerCmd.Flags().StringVarP(&connHost, "host", "s", "127.0.0.1", "host address")
	ServerCmd.Flags().StringVarP(&connPort, "port", "p", "7070", "host port")
	ServerCmd.Flags().StringVarP(&dbHost, "dbhost", "", "127.0.0.1", "db host")
	ServerCmd.Flags().StringVarP(&dbPort, "dbport", "", "5432", "db port")
	ServerCmd.Flags().StringVarP(&dbPassword, "dbpassword", "", "", "db password")
	ServerCmd.Flags().StringVarP(&dbName, "dbname", "n", "test", "db name")
	ServerCmd.Flags().StringVarP(&dbUser, "dbuser", "u", "", "db user")
	ServerCmd.Flags().StringVarP(&sslMode, "sslmode", "m", "disable", "ssl mode")
}

func serverStartCmdFunc(cmd *cobra.Command, args []string) {
	// Config-related part
	cfg, err := config.InitConfig(cfgPath)
	if err != nil {
		log.Fatalf("%s: InitConfig failed: %s", errors.ErrServiceCmdPrefix, err)
	}
	// Logger-related part
	logger, err := gcl.InitLogger(cfg.Level)
	if err != nil {
		log.Fatalf("%s: InitLogger failed: %s", errors.ErrServiceCmdPrefix, err)
	}
	// Init DB
	esp, err := dbFromConfig(cfg.DB)
	if err != nil {
		log.Fatalf("%s: dbFromConfig failed: %s", errors.ErrServiceCmdPrefix, err)
	}

	// Prometheus metrics start
	monitor, err := prometheus.NewMonitor(cfg.Prometheus)
	if err != nil {
		log.Fatalf("%s: Monitor init failed: %s", errors.ErrServiceCmdPrefix, err)
	}

	// Init GRPC server
	srv := grpc.NewServer(cfg, logger, esp, monitor)
	srv.Run()
}

func dbFromConfig(dbConfig config.DBConf) (interfaces.EventStorageProcessor, error) {
	if dbConfig.Name == "test" {
		inMemoryDB, err := db.NewInMemoryEventStorage()
		return inMemoryDB, err
	}
	mainDB, err := db.NewMainEventStorage(dbConfig)
	return mainDB, err
}
