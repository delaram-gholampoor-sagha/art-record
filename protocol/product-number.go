package protocol

// ProductNumber indicate the domain record data fields.
type ProductNumber interface {
	ProductID() [16]byte         // product domain
	Minimum() uint64             // Minimum number of sale to serve the service e.g. cinema wouldn't play movie, ...
	Maximum() uint64             // Maximum number can be served to buyers e.g. cinema tickets, panic-room game, ...
	MaximumInvoiceByBuy() uint64 // Maximum buy e.g. hotel room, panic-room game, ... that must limit by buyer.
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductNumber_StorageServices interface {
	Save(pa ProductNumber) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pa ProductNumber, numbers uint64, err protocol.Error)
}
