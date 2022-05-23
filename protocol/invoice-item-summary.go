package protocol

import "time"

// InvoiceItemSummary indicate the invoice-item-summary domain record data fields.
type InvoiceItemSummary interface {
	InvoiceID() [16]byte                     // invoice domain
	QuiddityID() [16]byte                    // product domain
	SuggestedPrice() protocol.AmountOfMoney  // Total should pay before any discount or auction
	DiscountedPrice() protocol.AmountOfMoney // user profit from this product purchase by auctions
	PayablePrice() protocol.AmountOfMoney    // Total price user pay to organization
	VAT() protocol.AmountOfMoney             // Total VAT user pay to its country government
	Time() time.Time                         // Save time
}

type InvoiceItemSummary_StorageServices interface {
	Save(iis InvoiceItemSummary) (err error)

	Count(invoiceID [16]byte) (numbers uint64, err error)
	Get(invoiceID [16]byte, versionOffset uint64) (iis InvoiceItemSummary, err error)
}
