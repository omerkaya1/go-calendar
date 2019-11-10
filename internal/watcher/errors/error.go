package errors

type WatcherError string

func (we WatcherError) Error() string {
	return string(we)
}

var (
	ErrBadConfigFile         = WatcherError("the correct configuration file was not specified")
	ErrBadQueueConfiguration = WatcherError("malformed or uninitialised message queue configuration")
)

const (
	ErrCMDPrefix = "command failure"
	ErrDBPrefix  = "db failure"
	ErrMQPrefix  = "message queue failure"
)
