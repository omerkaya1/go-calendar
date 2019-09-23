package grpc

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	cfgPath  string
	cfgType  string
	host     string
	port     string
	logLevel int
)

var ServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "Run GRPC Server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Implement me!")
		// TODO: there's a whole bunch of things, actually:
		//	 	 1) Init config from file, if it was supplied
		//	 	 2) Init config from cli, if 1) was not provided
	},
}

func init() {
	// NOTE: I'll have to decide how to properly handle server initialisation
	ServerCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	ServerCmd.PersistentFlags().IntVarP(&logLevel, "log", "l", 1, "changes log level")
	ServerCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "host address")
	ServerCmd.PersistentFlags().StringVarP(&port, "port", "p", "7070", "host port")
}
