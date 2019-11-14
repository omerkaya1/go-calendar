package errors

type NotificationError string

func (ne NotificationError) Error() string {
	return string(ne)
}

var (
	ErrBadConfigFile = NotificationError("the correct configuration file was not specified")
)

const (
	ErrCMDPrefix = "command failure"
	ErrMQPrefix  = "message queue failure"
)
