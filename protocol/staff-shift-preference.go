package protocol

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

// StaffShift indicate the domain record data fields.
type StaffShift interface {
	StaffID() [16]byte   // staff-status domain
	Day() utc.DayElapsed //
	RoleShift() uint8    //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
