package protocol

import (
	"../libgo/protocol"
)



type StaffEvaluation interface {
	StaffID() [16]byte   // staff-status domain
    EvaluatorID() [16]byte
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
