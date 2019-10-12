package parsers

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"log"
	"time"
)

// ParseTime .
func ParseTime(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		log.Fatalf("%s: %s", errors.ErrParsePrefix, "INVALID TIME!\n")
	}
	ts, err := ptypes.TimestampProto(*t)
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrParsePrefix, err)
	}
	return ts
}
