package protocol


// ProductReserveQuantity indicate the domain record data fields.
type ProductReserveQuantity interface {
	ProductID() [16]byte // product domain
	Quantity() uint64    // fix max reserve number
	Percent() uint64     // percent of the inventory
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductReserveQuantity_StorageServices interface {
	Save(pr ProductReserveQuantity) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pr ProductReserveQuantity, numbers uint64, err protocol.Error)
}
