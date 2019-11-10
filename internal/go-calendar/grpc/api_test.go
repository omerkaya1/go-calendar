package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/config"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/models"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/parsers"
	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/validators"
	gca "github.com/omerkaya1/go-calendar/internal/go-calendar/grpc/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"strings"
	"testing"
)

type testCalendar struct {
	client         gca.GoCalendarServerClient
	responseEvent  *models.Event
	updatedEvent   models.Event
	requestEvent   *models.EventJSON
	eventID        string
	updatedEventID string
	success        string
}

func (ct *testCalendar) _setup() error {
	cfg, err := config.InitConfig("./../../../configs/config.json")
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(cfg.Host+":"+cfg.Port, grpc.WithInsecure())
	if err != nil {
		return err
	}
	ct.client = gca.NewGoCalendarServerClient(conn)
	return nil
}

func (ct *testCalendar) everythingIsOk() error {
	return ct._setup()
}

func (ct *testCalendar) iMakeASendARequestToStoreAnEvent(data *gherkin.DocString) error {
	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(data.Content)
	event := &models.EventJSON{}
	if err := json.Unmarshal([]byte(cleanJson), event); err != nil {
		return err
	}
	ct.requestEvent = event
	start, err := validators.ValidateDate(ct.requestEvent.StartTime)
	if err != nil {
		return err
	}
	finish, err := validators.ValidateDate(ct.requestEvent.EndTime)
	if err != nil {
		return err
	}
	validators.ValidateTime(start, finish)
	startTime, err := parsers.ParseTimeToProto(start)
	if err != nil {
		return err
	}
	endTime, err := parsers.ParseTimeToProto(finish)
	if err != nil {
		return err
	}

	req := &gca.CreateEventRequest{
		UserName:  ct.requestEvent.UserName,
		EventName: ct.requestEvent.EventName,
		Text:      ct.requestEvent.Note,
		StartTime: startTime,
		EndTime:   endTime,
	}
	resp, err := ct.client.CreateEvent(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.GetError() != "" {
		return fmt.Errorf(resp.GetError())
	}
	if id := resp.GetEventID(); id != "" {
		ct.eventID = id
		return nil
	}
	return errors.ErrEmptyEventID
}

func (ct *testCalendar) iReceiveAnEventID() error {
	if ct.eventID != "" {
		return nil
	}
	return errors.ErrEmptyEventID
}

func (ct *testCalendar) iHaveTheEventID() error {
	return ct.iReceiveAnEventID()
}

func (ct *testCalendar) iRequestThisEventByItsID() error {
	req := &gca.RequestEventByID{
		EventID: ct.eventID,
	}
	var resp *gca.ResponseWithEvent
	var err error
	resp, err = ct.client.GetEvent(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.GetError() != "" {
		return fmt.Errorf(resp.GetError())
	}

	if e := resp.GetEvent(); e != nil {
		ct.responseEvent, err = parsers.MapProtoEventToEvent(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ct *testCalendar) theServerReturnsIt() error {
	if ct.responseEvent != nil {
		return nil
	}
	return errors.ErrEventDoesNotExist
}

func (ct *testCalendar) itMatchesTheTheOneWeSubmitted() error {
	var err error
	switch {
	case ct.responseEvent == nil:
		return errors.ErrEventDoesNotExist
	case ct.responseEvent.UserName != ct.requestEvent.UserName:
		err = fmt.Errorf("UserName fields don't match: %s and %s", ct.responseEvent.UserName, ct.requestEvent.UserName)
	case ct.responseEvent.EventName != ct.requestEvent.EventName:
		err = fmt.Errorf("EventName fields don't match: %s and %s", ct.responseEvent.EventName, ct.requestEvent.EventName)
	case ct.responseEvent.EventID.String() != ct.eventID:
		err = fmt.Errorf("EventID fields don't match: %s and %s", ct.responseEvent.EventID.String(), ct.requestEvent.EventID)
	case ct.responseEvent.Note != ct.requestEvent.Note:
		err = fmt.Errorf("the Note fields don't match: %s and %s", ct.responseEvent.Note, ct.requestEvent.Note)
		// TODO: Adjust for the time offset which has to be done because of the GRPC server behaviour
		//case ct.responseEvent.StartTime.Format(time.UnixDate) != ct.requestEvent.StartTime:
		//	err = fmt.Errorf("StartTime fields don't match: %s and %s", ct.responseEvent.StartTime.Format(time.UnixDate), ct.requestEvent.StartTime)
		//case ct.responseEvent.EndTime.Format(time.UnixDate) != ct.requestEvent.EndTime:
		//	err = fmt.Errorf("EndTime fields don't match: %s and %s", ct.responseEvent.EndTime.Format(time.UnixDate), ct.requestEvent.EndTime)
		break
	default:
		return nil
	}
	return err
}

func (ct *testCalendar) iUpdateTheStartTimeOfTheCreatedEventWithByItsID(data string) error {
	start, err := validators.ValidateDate(data)
	if err != nil {
		return err
	}
	ct.updatedEvent = *ct.responseEvent
	ct.updatedEvent.StartTime = start

	req := &gca.Event{
		EventName: ct.updatedEvent.EventName,
		Note:      ct.updatedEvent.Note,
	}

	pStart, err := parsers.ParseTimeToProto(start)
	if err != nil {
		return err
	}

	pFinish, err := parsers.ParseTimeToProto(ct.updatedEvent.EndTime)
	if err != nil {
		return err
	}

	req.StartTime, req.EndTime = pStart, pFinish
	req.EventId = ct.updatedEvent.EventID.String()

	resp, err := ct.client.UpdateEvent(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.GetError() != "" {
		return fmt.Errorf(resp.GetError())
	}
	ct.updatedEventID = resp.GetEventID()
	return nil
}

func (ct *testCalendar) theServerReturnsAnIDOfTheUpdatedEvent() error {
	if ct.updatedEventID != "" {
		return nil
	}
	return errors.ErrEmptyEventID
}

func (ct *testCalendar) bothIDsShouldMatch() error {
	if ct.updatedEventID != ct.eventID {
		return fmt.Errorf("EventID don't match: initial - %s updated - %s", ct.updatedEventID, ct.eventID)
	}
	return nil
}

func (ct *testCalendar) iRequestTheDeletionOfTheCreatedEventByItsID() error {
	req := &gca.RequestEventByID{
		EventID: ct.eventID,
	}
	resp, err := ct.client.DeleteEvent(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.GetError() != "" {
		return fmt.Errorf(resp.GetError())
	}
	ct.success = resp.GetResponse()
	return nil
}

func (ct *testCalendar) theServerReturnsASuccessMessage() error {
	if ct.success != "" {
		return nil
	}
	return fmt.Errorf("return message was empty")
}

func FeatureContext(s *godog.Suite) {
	test := testCalendar{}
	// Create a new event
	s.Step(`^everything is ok$`, test.everythingIsOk)
	s.Step(`^I make a send a request to store an event:$`, test.iMakeASendARequestToStoreAnEvent)
	s.Step(`^I receive an event ID$`, test.iReceiveAnEventID)
	// Get created event
	s.Step(`^I have the event ID$`, test.iHaveTheEventID)
	s.Step(`^I request this event by its ID$`, test.iRequestThisEventByItsID)
	s.Step(`^the server returns it$`, test.theServerReturnsIt)
	s.Step(`^it matches the the one we submitted$`, test.itMatchesTheTheOneWeSubmitted)
	// Update created event
	s.Step(`^I have the event ID$`, test.iHaveTheEventID)
	s.Step(`^I update the start time of the created event with "([^"]*)" by its ID$`, test.iUpdateTheStartTimeOfTheCreatedEventWithByItsID)
	s.Step(`^the server returns an ID of the updated event$`, test.theServerReturnsAnIDOfTheUpdatedEvent)
	s.Step(`^both IDs should match$`, test.bothIDsShouldMatch)
	// Delete created event
	s.Step(`^I have the event ID$`, test.iHaveTheEventID)
	s.Step(`^I request the deletion of the created event by its ID$`, test.iRequestTheDeletionOfTheCreatedEventByItsID)
	s.Step(`^the server returns a success message$`, test.theServerReturnsASuccessMessage)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGoDogAPI(t *testing.T) {
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"../../../test/integration/features/calendar"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})
	assert.Equal(t, 0, status)
}
