package protocol

// PosServiceArea indicate the domain record data fields.
// Save shortcut to regular products to easily add to invoice
type PosProduct interface {
	PosID() [16]byte     // pos domain
	ProductID() [16]byte // product domain.
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type PosProduct_StorageServices interface {
	Save(pp PosProduct) protocol.Error

	Count(posID [16]byte) (numbers uint64, err protocol.Error)
	Get(posID [16]byte, versionOffset uint64) (pp PosProduct, err protocol.Error)

	Delete(posID [16]byte, versionOffset uint64) (numbers uint64, err protocol.Error)
}
