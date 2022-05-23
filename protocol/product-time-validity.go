

package protocol

import (
	"../libgo/protocol"
)

type ProductTimeValidity interface {
	ProductID() [16]byte  // product-status domain
	Start() protocol.Time //
	End() protocol.Time   //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
