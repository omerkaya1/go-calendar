package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/omerkaya1/go-calendar/internal/integration-tests/calendar"
	"github.com/omerkaya1/go-calendar/internal/integration-tests/notification"
)

func main() {
	status := 1
	fmt.Println("Waiting for all services to become available...")
	// TODO: Create a basic script that will ping the PostgreSQL DB instead of using a sleep here
	time.Sleep(time.Second * 30)

	status = godog.RunWithOptions("calendar API", func(s *godog.Suite) {
		calendar.FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"./test/integration/features/calendar"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})

	status = godog.RunWithOptions("Notification", func(s *godog.Suite) {
		notification.FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"./test/integration/features/notification"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})
	os.Exit(status)
}
