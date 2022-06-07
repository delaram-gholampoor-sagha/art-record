package protocol

type ProductInventoryChange interface {
	ID() [16]byte        //
	ProductID() [16]byte // product domain
	From() [16]byte      // building-location domain
	To() [16]byte        // building-location domain. Can be zero for voiding products.
	Amount() int64       // this transaction
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
