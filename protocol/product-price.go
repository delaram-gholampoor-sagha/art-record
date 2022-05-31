package protocol

import (
	"../libgo/protocol"
)

// ProductPrice can be register just by producer organization whom own the related quiddity
type ProductPrice interface {
	ProductID() [16]byte           // product-status domain
	Currency() uint64              //
	Price() protocol.AmountOfMoney // Basic Price not payable price in the society currency
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}

type ProductPrice_StorageServices interface {
	Save(pp ProductPrice) (err protocol.Error)

	Count(productID [16]byte, currency uint64) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, currency uint64, versionOffset uint64) (pp ProductPrice, err protocol.Error)
	Last(productID [16]byte, currency uint64) (pp ProductPrice, numbers uint64, err protocol.Error)

	ListProductCurrencies(productID [16]byte, offset, limit uint64) (currencies []uint64, numbers uint64, err protocol.Error)
}
