package protocol

// ProductTag indicate the domain record data fields.
type ProductTag interface {
	ProductID() [16]byte       // product domain
	Tag() string               //
	Status() ProductTag_Status //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // user-request domain
}

type ProductTag_StorageServices interface {
	Save(pt ProductTag) protocol.Error

	Count(productID [16]byte, tag string) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, tag string, versionOffset uint64) (pt ProductTag, err protocol.Error)
	Last(productID [16]byte, tag string) (pt ProductTag, numbers uint64, err protocol.Error)

	ListTags(productID [16]byte, offset, limit uint64) (tags []string, numbers uint64, err protocol.Error)
	ListProducts(tag string, offset, limit uint64) (productIDs [16]byte, numbers uint64, err protocol.Error)

	FilterByStatus(productID [16]byte, status UserRelation_Status, offset, limit uint64) (tags []string, numbers uint64, err protocol.Error)
}

type ProductTag_Status Quiddity_Status

const (
// ProductTag_Status_ = ProductTag_Status(Quiddity_Status_FreeFlag << iota)
)
