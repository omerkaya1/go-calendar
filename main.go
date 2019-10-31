package main

import (
	"log"

	"github.com/omerkaya1/go-calendar/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
