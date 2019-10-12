package cmd

import (
	"github.com/omerkaya1/go-calendar/cmd/grpc"
	"github.com/omerkaya1/go-calendar/cmd/rws"

	"github.com/spf13/cobra"
)

// RootCmd only starts the app according to selected child command
var RootCmd = &cobra.Command{
	Use:   "go-calendar",
	Short: "simple calendar designed as a microservice",
}

func init() {
	RootCmd.AddCommand(grpc.ClientCmd, grpc.ServerCmd, rws.ClientCmd, rws.ServerCmd)
}
