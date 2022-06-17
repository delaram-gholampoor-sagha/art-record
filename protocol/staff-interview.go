package protocol

type StaffInterview interface {
	StaffID() [16]byte // staff domain

	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
