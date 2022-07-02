package protocol

// ProductContent indicate the domain record data fields.
type ProductContent interface {
	ProductID() [16]byte // product domain
	ContentID() [16]byte // content domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductContent_StorageServices interface {
	Save(pc ProductContent) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pc ProductContent, numbers uint64, err protocol.Error)

	FindByContent(contentID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}
