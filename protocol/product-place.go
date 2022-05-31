package protocol

import (
	"../libgo/protocol"
)

type ProductPlace interface {
	ProductID() [16]byte  // product-status domain.
	LocationID() [16]byte // location domain. use also for Coordinate domain
	MapID() [16]byte      // object domain. reserve seat?
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
