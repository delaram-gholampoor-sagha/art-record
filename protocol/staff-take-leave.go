package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type StaffTakeLeave interface {
	StaffID() [16]byte        // staff-status domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours //
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
