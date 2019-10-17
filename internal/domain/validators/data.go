package validators

import (
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

// ValidateDate checks the passed date string according to UNIX time format and returns address of the time object on
// success or logs the error on failure
func ValidateDate(timeStr string) (*time.Time, error) {
	t, err := time.Parse(time.UnixDate, timeStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// ValidateID checks the passed id string and returns the uuid object on success or logs the error on failure
func ValidateID(id string) (uuid.UUID, error) {
	verifiedID, err := uuid.FromString(id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return verifiedID, nil
}

func ValidateTime(start, finish *time.Time) {
	if start.After(*finish) {
		log.Fatalf("%s: %s", errors.ErrValidationPrefix, errors.ErrEventTimeViolation)
	}
}
