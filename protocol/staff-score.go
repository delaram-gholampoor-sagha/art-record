package protocol

// StaffScore or staff grade
type StaffScore interface {
	StaffID() [16]byte   // staff domain
	Score() uint64       //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
