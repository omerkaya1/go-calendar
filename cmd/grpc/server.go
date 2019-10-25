package grpc

import (
	"github.com/omerkaya1/go-calendar/internal/db"
	"github.com/omerkaya1/go-calendar/internal/domain/conf"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/services"
	"github.com/omerkaya1/go-calendar/internal/grpc"
	mq "github.com/omerkaya1/go-calendar/internal/message-queue"
	gcl "github.com/omerkaya1/go-calendar/log"
	"github.com/spf13/cobra"
	"log"
)

var (
	cfgPath  string
	connHost string
	connPort string
	dbName   string
	dbUser   string
	sslMode  string
	logLevel int
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
	ServerCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	ServerCmd.PersistentFlags().IntVarP(&logLevel, "log", "l", 1, "changes log level")
	ServerCmd.PersistentFlags().StringVarP(&connHost, "host", "s", "127.0.0.1", "host address")
	ServerCmd.PersistentFlags().StringVarP(&connPort, "port", "p", "7070", "host port")
	ServerCmd.PersistentFlags().StringVarP(&dbName, "dbname", "n", "test", "db name")
	ServerCmd.PersistentFlags().StringVarP(&dbUser, "dbuser", "u", "", "db user")
	ServerCmd.PersistentFlags().StringVarP(&sslMode, "sslmode", "m", "disable", "ssl mode")
}

// TODO: consider using it as a separate method, checking the command's name and calling the corresponding server
func serverStartCmdFunc(cmd *cobra.Command, args []string) {
	// Config-related part
	cfg := &conf.Config{}
	var err error
	if cfgPath != "" {
		cfg, err = conf.CfgFromFile(cfgPath)
	} else {
		cfg = conf.CfgFromCmdParams(logLevel, connHost, connPort, dbName, dbUser, sslMode)
	}
	if err != nil {
		log.Fatalf("%s: InitConfig failed: %s", errors.ErrServiceCmdPrefix, err)
	}
	// Logger-related part
	logger, err := gcl.InitLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("%s: InitLogger failed: %s", errors.ErrServiceCmdPrefix, err)
	}
	// Init DB
	esp, err := dbFromConfig(cfg.DB)
	if err != nil {
		log.Fatalf("%s: dbFromConfig failed: %s", errors.ErrServiceCmdPrefix, err)
	}

	// Init MQ
	q, err := mq.NewMessageQueue(cfg.Queue)
	if err != nil {
		log.Fatalf("%s: NewMessageQueue failed: %s", errors.ErrServiceCmdPrefix, err)
	}

	// Init GRPC server
	srv := grpc.NewServer(cfg, logger, esp, q)
	srv.Run()
}

// TODO: This should be moved somewhere as well
func dbFromConfig(dbConfig conf.DBConf) (*services.EventService, error) {
	if dbConfig.Name == "test" {
		inMemoryDB, err := db.NewInMemoryEventStorage()
		return &services.EventService{Processor: inMemoryDB}, err
	}
	mainDB, err := db.NewMainEventStorage(dbConfig)
	return &services.EventService{Processor: mainDB}, err
}
