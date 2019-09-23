package main

import (
	"fmt"
	"log"
	"os"
)

var timeout int

var rootCmd = &cobra.Command{
	Use:   "go-calendar [HOST ADDRESS] [PORT]",
	Short: "simple calendar designed as a microservice",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("invalid number of arguments")
		}
		return nil
	},
	Run: rootCommand,
}

// Execute is a method that runs the root command of the programme
func Execute() {
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 30, "timeout before exiting the programme (default is 30s")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCommand(cmd *cobra.Command, args []string) {
	s := internal.NewNetworkConn(timeout, args[0], args[1])
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
