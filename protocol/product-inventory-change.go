package protocol

import (
	"../libgo/protocol"
)

type ProductInventoryChange interface {
	ID() [16]byte             //
	ProductID() [16]byte      // product-status domain
	FromLocationID() [16]byte // location domain
	ToLocationID() [16]byte   // location domain. Can be zero for voiding products.
	Amount() int64            // this transaction
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
