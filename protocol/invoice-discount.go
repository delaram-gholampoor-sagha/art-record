package protocol

// InvoiceDiscount indicate the domain record data fields.
type InvoiceDiscount interface {
	InvoiceID() [16]byte                // invoice domain
	DiscountID() [16]byte               // discount domain
	Discounted() protocol.AmountOfMoney //
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}

type InvoiceDiscount_StorageServices interface {
	Save(iv InvoiceDiscount) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iv InvoiceDiscount, err protocol.Error)
	Last(invoiceID [16]byte) (iv InvoiceDiscount, numbers uint64, err protocol.Error)

	FindByDiscountID(discountID [16]byte, offset, limit uint64) (invoiceIDs [][16]byte, numbers uint64, err protocol.Error)
}
