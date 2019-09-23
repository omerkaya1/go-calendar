package errors

type GoCalendarError string

func (ee GoCalendarError) Error() string {
	return string(ee)
}

var (
	ErrCorruptConfigFileExtension = GoCalendarError("another event exists for this date")
	ErrEventCollision             = GoCalendarError("another event exists for this date")
	//ErrIncorrectEndDate = GoCalendarError("end_date is incorrect")
)
