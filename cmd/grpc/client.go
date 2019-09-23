package grpc

import (
	"github.com/spf13/cobra"
	"log"
)

var ClientCmd = &cobra.Command{
	Use:   "grpc-client",
	Short: "Run GRPC client",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Implement me!")
	},
}

func init() {
	// Init some flags here
}
