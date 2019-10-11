package rws

import (
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/spf13/cobra"
	"log"
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
		-s "Tue Oct 1 18:00:00 MSK 2019" -e "Tue Oct 1 23:30:00 MSK 2019"`,
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
-s "Tue Oct 1 19:00:00 MSK 2019" -e "Tue Oct 1 23:30:00 MSK 2019"`,
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
	ClientCmd.PersistentFlags().StringVarP(&port, "port", "p", "9000", "port of the host")
	ClientCmd.PersistentFlags().StringVarP(&eventID, "id", "i", "", "internal event id")
	ClientCmd.PersistentFlags().StringVarP(&eventOwner, "owner", "o", "", "owner of the event")
	ClientCmd.PersistentFlags().StringVarP(&eventName, "event-title", "t", "", "event name")
	ClientCmd.PersistentFlags().StringVarP(&eventNote, "note", "n", "empty", "additional note related to the event")
	ClientCmd.PersistentFlags().StringVarP(&startTime, "event-start", "b", "", "starting date and hour of the event")
	ClientCmd.PersistentFlags().StringVarP(&endTime, "event-end", "e", "", "ending date and hour of the event")
}

// MEMO: consider using gorilla Client
func createCmdFunc(cmd *cobra.Command, args []string) {
	if eventOwner == "" || startTime == "" || endTime == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrUnsetFlags.Error())
	}
	//start := validators.ValidateDate(startTime)
	//finish := validators.ValidateDate(endTime)
	//
	//eventID, err :=
}

func updateCmdFunc(cmd *cobra.Command, args []string) {
	log.Println("Implement me!")
	log.Println(eventName, eventID, "|", eventNote, "|", eventOwner, startTime, endTime)
}

func getCmdFunc(cmd *cobra.Command, args []string) {
	log.Println("Implement me!")
	log.Println(eventName, eventID, "|", eventNote, "|", eventOwner, startTime, endTime)
}

func deleteCmdFunc(cmd *cobra.Command, args []string) {
	log.Println("Implement me!")
	log.Println(eventName, eventID, "|", eventNote, "|", eventOwner, startTime, endTime)
}
