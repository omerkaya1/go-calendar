package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/omerkaya1/go-calendar/internal/db"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/validators"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type testNotification struct {
	n            *RabbitMQService
	eventToStore *models.Event
	requestEvent *models.EventJSON
	eventID      string
}

func (tn *testNotification) _setup() error {
	var err error
	conf, err := config.InitConfig("./../../configs/config.json")
	if err != nil {
		return err
	}

	esp, err := db.NewMainEventStorage(conf.DB)
	if err != nil {
		return err
	}
	tn.n, err = NewRabbitMQService(conf, esp)
	return err
}

func (tn *testNotification) aNewEventIsStoredInTheDB(data *gherkin.DocString) error {
	if err := tn._setup(); err != nil {
		return err
	}
	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(data.Content)
	event := &models.EventJSON{}
	if err := json.Unmarshal([]byte(cleanJson), event); err != nil {
		return err
	}
	tn.requestEvent = event
	start, err := validators.ValidateDate(tn.requestEvent.StartTime)
	if err != nil {
		return err
	}
	finish, err := validators.ValidateDate(tn.requestEvent.EndTime)
	if err != nil {
		return err
	}
	validators.ValidateTime(start, finish)

	tn.eventToStore = models.NewEvent(event.UserName, event.EventName, event.Note, start, finish)

	resp, err := tn.n.db.CreateEvent(context.Background(), tn.eventToStore)
	if err != nil {
		return err
	}
	if resp.String() != "" {
		return nil
	}
	return errors.ErrEmptyEventID
}

func (tn *testNotification) theNotificationServiceIsStarted() error {
	go func() {
		if err := tn.n.ProduceMessages(); err != nil {
			return
		}
	}()
	return nil
}

func (tn *testNotification) theNotificationServiceReturnsCreatedEventAsAMessage() error {
	ch, err := tn.n.Conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		"events",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	retVal := ""
	for m := range msgs {
		retVal = string(m.Body)
		break
	}
	switch {
	case !strings.Contains(retVal, tn.requestEvent.UserName):
	case !strings.Contains(retVal, tn.requestEvent.EventName):
	case !strings.Contains(retVal, tn.requestEvent.StartTime):
	case !strings.Contains(retVal, tn.requestEvent.EndTime):
		return fmt.Errorf("some shit")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := testNotification{}
	s.Step(`^a new event is stored in the DB:$`, test.aNewEventIsStoredInTheDB)
	s.Step(`^the notification service is started$`, test.theNotificationServiceIsStarted)
	s.Step(`^the notification service returns created event as a message$`, test.theNotificationServiceReturnsCreatedEventAsAMessage)
}

func TestGoDogNotification(t *testing.T) {
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"../../test/integration/features/notification"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})
	assert.Equal(t, 0, status)
}
