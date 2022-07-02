package protocol

// InvoiceItemSummary indicate the domain record data fields.
type InvoiceItemSummary interface {
	InvoiceID() [16]byte                     // invoice domain
	ProductID() [16]byte                     // product domain
	SuggestedPrice() protocol.AmountOfMoney  // Total should pay before any discount or auction
	DiscountedPrice() protocol.AmountOfMoney // user profit from this product purchase by auctions
	PayablePrice() protocol.AmountOfMoney    // Total price user pay to organization
	VAT() protocol.AmountOfMoney             // Total VAT user pay to its country government
	Time() protocol.Time                     // save time
}

type InvoiceItemSummary_StorageServices interface {
	Save(iis InvoiceItemSummary) (numbers uint64, err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iis InvoiceItemSummary, err protocol.Error)
}
