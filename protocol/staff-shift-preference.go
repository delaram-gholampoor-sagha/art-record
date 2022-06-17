package protocol


// StaffShift indicate the domain record data fields.
type StaffShift interface {
	StaffID() [16]byte   // staff domain
	Day() utc.DayElapsed //
	ShiftID() [16]byte   // org-shift domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
