package main

import (
	"github.com/omerkaya1/go-calendar/cmd"
	"log"
)

//go:generate echo test

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
