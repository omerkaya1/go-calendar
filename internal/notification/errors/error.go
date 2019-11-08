package errors

type NotificationError string

func (ne NotificationError) Error() string {
	return string(ne)
}

var (
	ErrBadConfigFile         = NotificationError("the correct configuration file was not specified")
	ErrBadQueueConfiguration = NotificationError("malformed or uninitialised message queue configuration")
)

const (
	ErrCMDPrefix = "command failure"
	ErrMQPrefix  = "message queue failure"
)
