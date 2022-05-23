package protocol

import (
	"../libgo/protocol"
)

type CategoryComplement interface {
	CategoryID() [16]byte   // category domain
	Priotry() int32         // Complement priotry
	ComplementID() [16]byte // category domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}
