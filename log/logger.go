package log

import (
	"log"

	"go.uber.org/zap"
)

// InitLogger method sets up a server logger according to the specified log level
func InitLogger(level int) (*zap.Logger, error) {
	l := &zap.Logger{}
	var err error
	switch level {
	case 0:
		l = zap.NewExample()
	case 1:
		l, err = zap.NewProduction()
	case 2:
		l, err = zap.NewDevelopment()
	default:
		log.Panicf("incorrect logging level: %v", level)
	}
	return l, err
}
