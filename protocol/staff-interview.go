package protocol

import "../libgo/protocol"

type StaffInterview interface {
	StaffID() [16]byte // staff-status domain

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
