package protocol

import (
	"../libgo/protocol"
)

type ProductSubstitute interface {
	ProductID() [16]byte    // product-status domain
	Priotry() uint64        // Substitute priotry
	SubstituteID() [16]byte // product-status domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}
