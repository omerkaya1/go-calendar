package rws

import (
	"github.com/spf13/cobra"
	"log"
)

var ClientCmd = &cobra.Command{
	Use:   "rws-client",
	Short: "Run RESTful Web Service client",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Implement me!")
	},
}

func init() {
	// Init some flags here
}
