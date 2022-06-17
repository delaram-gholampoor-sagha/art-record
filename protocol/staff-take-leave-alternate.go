package protocol


// StaffTakeLeaveAlternate or staff replace
type StaffTakeLeaveAlternate interface {
	StaffID() [16]byte     // staff domain
	Day() utc.DayElapsed   //
	PresentedID() [16]byte // staff domain
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}
