package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)
// the time that a satff has to be in a role ?
// برنامه ی کاری یک استف با یک رول و یک دیپارتمان متفاوت است
type StaffTimingUTC interface {
	StaffID() [16]byte        // staff-status domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours // what hours in a day this role active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}

