package notification

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/omerkaya1/go-calendar/internal/db"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"
	"github.com/omerkaya1/go-calendar/internal/mq"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
	"time"
)

type testNotification struct {
	n            mq.MessageQueue
	eventToStore *models.Event
	eventID      uuid.UUID
}

func setupNotification() *testNotification {
	var err error
	conf, err := config.InitConfig("./configs/config.json")
	if err != nil {
		log.Fatal(err)
	}

	mdb, err := db.NewMainEventStorage(conf.DB)
	if err != nil {
		log.Fatal(err)
	}
	s := time.Now()
	f := time.Now().Add(3 * time.Hour)
	e := models.NewEvent("Rick", "party", "get schwifty", &s, &f)

	id, err := mdb.CreateEvent(context.Background(), e)
	if err != nil {
		log.Fatal(err)
	}

	rmq, err := mq.NewRabbitMQService(conf, mdb, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &testNotification{
		n:            rmq,
		eventToStore: e,
		eventID:      id,
	}
}

func (tn *testNotification) everythingIsSetUp() error {
	if tn.eventToStore == nil && tn.eventID.String() == "" {
		return fmt.Errorf("setup failed")
	}
	return nil
}

func (tn *testNotification) anEventIsStoredInTheDB() error {
	if tn.eventID.String() != "" {
		return nil
	}
	return errors.ErrEmptyEventID
}

func (tn *testNotification) theNotificationServiceIsStarted() error {
	go func() {
		if err := tn.n.ProduceMessages(); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (tn *testNotification) theNotificationServiceReturnsStoredEventAsAMessage() error {
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
	case !strings.Contains(retVal, tn.eventToStore.UserName):
	case !strings.Contains(retVal, tn.eventToStore.EventName):
	case !strings.Contains(retVal, tn.eventToStore.StartTime.String()):
	case !strings.Contains(retVal, tn.eventToStore.EndTime.String()):
		return fmt.Errorf("the returned message doesn't include requered fields")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := setupNotification()
	s.Step(`^everything is set up$`, test.everythingIsSetUp)
	s.Step(`^an event is stored in the DB$`, test.anEventIsStoredInTheDB)
	s.Step(`^the notification service is started$`, test.theNotificationServiceIsStarted)
	s.Step(`^the notification service returns stored event as a message$`, test.theNotificationServiceReturnsStoredEventAsAMessage)
}
