package protocol

// ProductGroup indicate the domain record data fields.
type ProductGroup interface {
	ProductID() [16]byte // product domain
	GroupID() [16]byte   // group domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductGroup_StorageServices interface {
	Save(pg ProductGroup) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pg ProductGroup, numbers uint64, err protocol.Error)

	FindByGroup(groupID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}
