package errors

type GoCalendarError string

func (ee GoCalendarError) Error() string {
	return string(ee)
}

var (
	// CMD related errors
	ErrCorruptConfigFileExtension = GoCalendarError("file extension was not determined")
	ErrBadDBConfiguration         = GoCalendarError("malformed or uninitialised DB configuration")
	ErrUnsetFlags                 = GoCalendarError("some flags are missing or unset")
	// Event conflict related errors
	ErrEventCollisionInInterval = GoCalendarError("event takes place within the time interval of another event")
	ErrEventDoesNotExist        = GoCalendarError("the requested event does not exist in the DB")
	ErrEventTimeViolation       = GoCalendarError("new events cannot be created in the past")
	ErrMalformedTimeObject      = GoCalendarError("invalid time string")
	// DB related errors
	ErrNoOpDBAction = GoCalendarError("no rows were affected by the action")
)

const (
	ErrServiceCmdPrefix = "server failure"
	ErrClientCmdPrefix  = "client failure"
	ErrAPIPrefix        = "api failure"
	ErrValidationPrefix = "validation failure"
)
