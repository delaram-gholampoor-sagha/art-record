package protocol

// ProductComplement indicate the domain record data fields.
type ProductComplement interface {
	ProductID() [16]byte    // product domain
	Priority() uint64       // Complement priority
	ComplementID() [16]byte // product domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductComplement_StorageServices interface {
	Save(pc ProductComplement) (numbers uint64, err protocol.Error)

	Count(productID [16]byte, priority uint64) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, priority uint64, versionOffset uint64) (pc ProductComplement, numbers uint64, err protocol.Error)

	ListProducts(complementID [16]byte, offset, limit uint64) (productIDs [16]byte, numbers uint64, err protocol.Error)
	ListPriorities(complementID [16]byte, productID [16]byte, offset, limit uint64) (priorities []uint64, numbers uint64, err protocol.Error)
}
