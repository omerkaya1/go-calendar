package errors

type GoCalendarError string

func (ee GoCalendarError) Error() string {
	return string(ee)
}

var (
	ErrCorruptConfigFileExtension = GoCalendarError("file extension was not determined")
	ErrBadDBConfiguration         = GoCalendarError("malformed or uninitialised DB configuration")
	ErrUnsetFlags                 = GoCalendarError("some flags are missing or unset")
	ErrEventCollisionInInterval   = GoCalendarError("event takes place within the time interval of another event")
	ErrEventCollisionOutInterval  = GoCalendarError("another event takes place within the new event's timeline")
	ErrEventCollisionBefore       = GoCalendarError("new event ends in the middle of another's event timeline")
	ErrEventCollisionAfter        = GoCalendarError("new event starts in the middle of another's event timeline")
	ErrEventCollisionMatch        = GoCalendarError("new event cannot take place at the same time with another event")
	ErrEventDoesNotExist          = GoCalendarError("the requested event does not exist in the DB")
	ErrEventTimeViolation         = GoCalendarError("new events cannot be created in the past")
)

const (
	ErrServiceCmdPrefix = "server failure"
	ErrClientCmdPrefix  = "client failure"
	ErrValidationPrefix = "validation failure"
	ErrParsePrefix      = "parse failure"
)
