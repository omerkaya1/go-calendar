package main

import (
	"github.com/omerkaya1/go-calendar/cmd/go-calendar/grpc"
	"github.com/omerkaya1/go-calendar/cmd/go-calendar/rws"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-calendar",
	Short: "simple calendar designed as a microservice",
}

func init() {
	rootCmd.AddCommand(grpc.ClientCmd, grpc.ServerCmd, rws.ClientCmd, rws.ServerCmd)
}
