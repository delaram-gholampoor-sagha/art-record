package protocol

type StaffOvertime interface {
	StaffID() [16]byte        // staff domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}
