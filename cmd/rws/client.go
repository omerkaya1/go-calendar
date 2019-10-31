package rws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/models"
	"github.com/omerkaya1/go-calendar/internal/rws"
	"github.com/spf13/cobra"
)

var host, port, eventName, eventID, eventNote, eventOwner, startTime, endTime string

var (
	ClientCmd = &cobra.Command{
		Use:     "rws-client",
		Short:   "Run RESTful Web Service client",
		Example: "  go-calendar rws-client create -h",
	}

	CreateActionCmd = &cobra.Command{
		Use:   "create",
		Short: "Create calendar event",
		Run:   createCmdFunc,
		Example: `  go-calendar rws-client create -t "Saturday party" -n "Buy soda and apples!" -o "John Doe" 
		-b "Tue Oct 1 18:00:00 MSK 2019" -e "Tue Oct 1 23:30:00 MSK 2019"`,
	}

	GetActionCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get calendar event",
		Run:     getCmdFunc,
		Example: "  go-calendar rws-client get -i sdkjf-8783-sdfs-341\n  go-calendar rws-client -o \"John Doe\"",
	}

	UpdateActionCmd = &cobra.Command{
		Use:   "update",
		Short: "Update calendar event",
		Run:   updateCmdFunc,
		Example: `  go-calendar rws-client update -i sdkjf-8783-sdfs-341 -t "Saturday party(postponed)" -o "John Doe" 
-b "Tue Oct 1 19:00:00 MSK 2019" -e "Tue Oct 1 23:30:00 MSK 2019"`,
	}

	DeleteActionCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete calendar event",
		Run:   deleteCmdFunc,
		Example: "  go-calendar rws-client delete -i sdkjf-8783-sdfs-341\n" +
			"  go-calendar rws-client delete -t \"Saturday party(postponed)\"" +
			"  go-calendar rws-client delete -o \"John Doe\"",
	}
)

func init() {
	ClientCmd.AddCommand(CreateActionCmd, GetActionCmd, UpdateActionCmd, DeleteActionCmd)
	ClientCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "host address to connect to")
	ClientCmd.PersistentFlags().StringVarP(&port, "port", "p", "7070", "port of the host")
	ClientCmd.PersistentFlags().StringVarP(&eventID, "id", "i", "", "internal event id")
	ClientCmd.PersistentFlags().StringVarP(&eventOwner, "owner", "o", "", "owner of the event")
	ClientCmd.PersistentFlags().StringVarP(&eventName, "event-title", "t", "", "event name")
	ClientCmd.PersistentFlags().StringVarP(&eventNote, "note", "n", "empty", "additional note related to the event")
	ClientCmd.PersistentFlags().StringVarP(&startTime, "event-start", "b", "", "starting date and hour of the event")
	ClientCmd.PersistentFlags().StringVarP(&endTime, "event-end", "e", "", "ending date and hour of the event")
}

func createCmdFunc(cmd *cobra.Command, args []string) {
	if eventOwner == "" || startTime == "" || endTime == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getRWSClient()

	event := &models.EventJSON{
		EventID:   eventID,
		UserName:  eventOwner,
		EventName: eventName,
		Note:      eventNote,
		StartTime: startTime,
		EndTime:   endTime,
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://%s:%s%s%s%s", host, port, rws.RWSApiPrefix, rws.RWSapiVersion, rws.RWSeventURL),
		bytes.NewReader(body))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	defer resp.Body.Close()
	buf := make([]byte, resp.ContentLength)
	if n, err := resp.Body.Read(buf); err.Error() != "EOF" || int64(n) != resp.ContentLength {
		log.Fatalf("%s: %s. bytes read: %d", errors.ErrClientCmdPrefix, err, n)
	}
	log.Println(string(buf))
}

func updateCmdFunc(cmd *cobra.Command, args []string) {
	if eventID == "" && eventName == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getRWSClient()

	event := &models.EventJSON{EventID: eventID, UserName: eventOwner, EventName: eventName, Note: eventNote}
	if startTime != "" && endTime != "" {
		event.StartTime, event.EndTime = startTime, endTime
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://%s:%s%s%s%s", host, port, rws.RWSApiPrefix, rws.RWSapiVersion, rws.RWSeventURL),
		bytes.NewReader(body))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	defer resp.Body.Close()
	buf := make([]byte, resp.ContentLength)
	if n, err := resp.Body.Read(buf); err.Error() != "EOF" || int64(n) != resp.ContentLength {
		log.Fatalf("%s: %s. bytes read: %d", errors.ErrClientCmdPrefix, err, n)
	}
	log.Println(string(buf))
}

func getCmdFunc(cmd *cobra.Command, args []string) {
	if eventID == "" && eventName == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getRWSClient()

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://%s:%s%s%s%s/%s", host, port, rws.RWSApiPrefix, rws.RWSapiVersion, rws.RWSeventURL, eventID),
		nil)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	defer resp.Body.Close()
	buf := make([]byte, resp.ContentLength)
	if n, err := resp.Body.Read(buf); err.Error() != "EOF" || int64(n) != resp.ContentLength {
		log.Fatalf("%s: %s. bytes read: %d", errors.ErrClientCmdPrefix, err, n)
	}
	log.Println(string(buf))
}

func deleteCmdFunc(cmd *cobra.Command, args []string) {
	if eventID == "" && eventName == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getRWSClient()

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("http://%s:%s%s%s%s/%s", host, port, rws.RWSApiPrefix, rws.RWSapiVersion, rws.RWSeventURL, eventID),
		nil)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	defer resp.Body.Close()
	buf := make([]byte, resp.ContentLength)
	if n, err := resp.Body.Read(buf); err.Error() != "EOF" || int64(n) != resp.ContentLength {
		log.Fatalf("%s: %s. bytes read: %d", errors.ErrClientCmdPrefix, err, n)
	}
	log.Println(string(buf))
}

func getRWSClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
	}
}
