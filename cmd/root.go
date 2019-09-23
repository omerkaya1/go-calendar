package cmd

import (
	//"log"
	"os"

	"github.com/spf13/cobra"
)

var cfgPath string

var rootCmd = &cobra.Command{
	Use:   "go-calendar",
	Short: "simple calendar designed as a microservice",
	Run:   rootCommand,
}

// Execute is a method that runs the root command of the programme
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "path to the configuration file")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCommand(cmd *cobra.Command, args []string) {
	//s := internal.NewNetworkConn(timeout, args[0], args[1])
	//if err := s.Serve(); err != nil {
	//	log.Fatal(err)
	//}
}
