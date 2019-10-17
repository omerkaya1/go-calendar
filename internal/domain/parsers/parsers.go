package parsers

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"time"
)

// ParseTime .
func ParseTimeToProto(t *time.Time) (*timestamp.Timestamp, error) {
	if t == nil {
		return nil, errors.ErrMalformedTimeObject
	}
	ts, err := ptypes.TimestampProto(*t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func ParseProtoToTime(t *timestamp.Timestamp) (*time.Time, error) {
	if t == nil {
		return nil, errors.ErrMalformedTimeObject
	}
	st, err := ptypes.Timestamp(t)
	if err != nil {
		return nil, err
	}
	return &st, nil
}
