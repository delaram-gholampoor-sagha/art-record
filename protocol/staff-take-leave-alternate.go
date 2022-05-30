package protocol

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

// the person that is your alternate 

// StaffTakeLeaveAlternate or staff replace
type StaffTakeLeaveAlternate interface {
	StaffID() [16]byte     // staff-status domain
	Day() utc.DayElapsed   //
	PresentedID() [16]byte // staff-status domain
	Time() protocol.Time   // Save time
	RequestID() [16]byte   // user-request domain
}
