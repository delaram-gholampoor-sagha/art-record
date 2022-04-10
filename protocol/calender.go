package protocol

import "time"

// we need the time concept in out software (sth like calender)
type Calender struct {
	ID          uint64
	QuiddityID  [16]byte
	Status      uint8
	CurrentDate time.Time
	// SpecifiedDate time.Time
	// ElapsedTime ??
	Created_At time.Time
	Updated_At time.Time
}

type CalenderServices interface {
	GetCalenderState(CalenderID uint64) (string, error)
	UpdateCalenderDate(isPartial bool, calender Calender) (Calender, error)
}
