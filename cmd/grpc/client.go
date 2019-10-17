package grpc

import (
	"context"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/omerkaya1/go-calendar/internal/domain/parsers"
	"github.com/omerkaya1/go-calendar/internal/domain/validators"
	gca "github.com/omerkaya1/go-calendar/internal/grpc/go-calendar-api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

var host, port, eventName, eventID, eventNote, eventOwner, startTime, endTime string

var (
	ClientCmd = &cobra.Command{
		Use:     "grpc-client",
		Short:   "Run GRPC Web Service client",
		Example: "  go-calendar grpc-client create -h",
	}

	CreateActionCmd = &cobra.Command{
		Use:   "create",
		Short: "Create calendar event",
		Run:   createCmdFunc,
		Example: `  go-calendar grpc-client create -t "Saturday party" -n "Buy soda and apples!" -o "John Doe" 
		-s "Tue Oct 1 18:00:00 MSK 2019" -e "Tue Oct 1 23:30:00 MSK 2019"`,
	}

	GetActionCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get calendar event",
		Run:     getCmdFunc,
		Example: "  go-calendar grpc-client get -i sdkjf-8783-sdfs-341\n  go-calendar grpc-client -o \"John Doe\"",
	}

	UpdateActionCmd = &cobra.Command{
		Use:   "update",
		Short: "Update calendar event",
		Run:   updateCmdFunc,
		Example: `  go-calendar grpc-client update -i sdkjf-8783-sdfs-341 -t "Saturday party(postponed)" -o "John Doe" 
-s "Tue Oct 1 19:00:00 MSK 2019" -e "Tue Oct 1 23:30:00 MSK 2019"`,
	}

	DeleteActionCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete calendar event",
		Run:   deleteCmdFunc,
		Example: "  go-calendar grpc-client delete -i sdkjf-8783-sdfs-341\n" +
			"  go-calendar grpc-client delete -t \"Saturday party(postponed)\"" +
			"  go-calendar grpc-client delete -o \"John Doe\"",
	}
)

func init() {
	ClientCmd.AddCommand(CreateActionCmd, GetActionCmd, UpdateActionCmd, DeleteActionCmd)
	ClientCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "host address to connect to")
	ClientCmd.PersistentFlags().StringVarP(&port, "port", "p", "7070", "port of the host")
	ClientCmd.PersistentFlags().StringVarP(&eventID, "id", "i", "", "internal event id")
	ClientCmd.PersistentFlags().StringVarP(&eventOwner, "owner", "o", "", "owner of the event")
	ClientCmd.PersistentFlags().StringVarP(&eventName, "event-title", "t", "", "event name")
	ClientCmd.PersistentFlags().StringVarP(&eventNote, "note", "n", "", "additional note related to the event")
	ClientCmd.PersistentFlags().StringVarP(&startTime, "event-start", "b", "", "starting date and hour of the event")
	ClientCmd.PersistentFlags().StringVarP(&endTime, "event-end", "e", "", "ending date and hour of the event")
}

func createCmdFunc(cmd *cobra.Command, args []string) {
	if eventOwner == "" || startTime == "" || endTime == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getGRPCClient()

	start, err := validators.ValidateDate(startTime)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	finish, err := validators.ValidateDate(endTime)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	validators.ValidateTime(start, finish)

	startTime, err := parsers.ParseTimeToProto(start)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	endTime, err := parsers.ParseTimeToProto(finish)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	req := &gca.CreateEventRequest{
		UserName:  eventOwner,
		EventName: eventName,
		Text:      eventNote,
		StartTime: startTime,
		EndTime:   endTime,
	}
	resp, err := client.CreateEvent(context.Background(), req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Println(resp.GetEventID())
}

func updateCmdFunc(cmd *cobra.Command, args []string) {
	if eventID == "" && eventName == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getGRPCClient()

	req := &gca.Event{
		EventName: eventName,
		Note:      eventNote,
	}

	start, finish := &time.Time{}, &time.Time{}
	if startTime != "" && endTime != "" {
		start, err := validators.ValidateDate(startTime)
		if err != nil {
			log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
		}
		finish, err := validators.ValidateDate(endTime)
		if err != nil {
			log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
		}
		validators.ValidateTime(start, finish)
	}

	pStart, err := parsers.ParseTimeToProto(start)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	pFinish, err := parsers.ParseTimeToProto(finish)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	req.StartTime, req.EndTime = pStart, pFinish

	id, err := validators.ValidateID(eventID)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	req.EventId = id.String()

	resp, err := client.UpdateEvent(context.Background(), req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Println(resp.GetEventID())
}

func getCmdFunc(cmd *cobra.Command, args []string) {
	if eventID == "" && eventName == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getGRPCClient()

	id, err := validators.ValidateID(eventID)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	req := &gca.RequestEventByID{
		EventID: id.String(),
	}

	resp, err := client.GetEvent(context.Background(), req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Printf("%v\n", resp.GetEvent())
}

func deleteCmdFunc(cmd *cobra.Command, args []string) {
	if eventID == "" && eventName == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags)
	}

	client := getGRPCClient()

	id, err := validators.ValidateID(eventID)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}

	req := &gca.RequestEventByID{
		EventID: id.String(),
	}

	resp, err := client.DeleteEvent(context.Background(), req)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Printf("%v\n", resp.GetResponse())
}

func getGRPCClient() gca.GoCalendarServerClient {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	return gca.NewGoCalendarServerClient(conn)
}
