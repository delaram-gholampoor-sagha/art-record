package protocol

import (
	"../libgo/protocol"
)

// ProductPricePayable store historically price with discount to show to users in price change chart.
type ProductPricePayable interface {
	ProductID() [16]byte // product-status domain
	// OrgID() [16]byte               //
	Currency() uint64              //
	Price() protocol.AmountOfMoney // Basic Price not payable price in the society currency
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}

type ProductPrice_StorageServices interface {
	Save(q ProductPricePayable) (err protocol.Error)

	Count(productID [16]byte, currency uint64) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, currency uint64, versionOffset uint64) (q ProductPricePayable, err protocol.Error)
	Last(productID [16]byte, currency uint64) (q ProductPricePayable, numbers uint64, err protocol.Error)
}
