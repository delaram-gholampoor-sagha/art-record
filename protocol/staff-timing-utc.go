package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type StaffTimingUTC interface {
	StaffID() [16]byte        // staff-status domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours // what hours in a day this role active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
