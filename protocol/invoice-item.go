package protocol

import "time"

// InvoiceItem indicate the invoice-item domain record data fields.
// each InvoiceItem is immutable record and so use version mechanism to chain InvoiceItem in InvoiceID group.
type InvoiceItem interface {
	InvoiceID() [16]byte  // invoice domain
	QuiddityID() [16]byte // product domain
	Quantity() uint32     // decimal >> 1.5 >> 1.2 >> 10 >> 0
	Time() time.Time      // Save time
	RequestID() [16]byte  // user-request domain
}

type InvoiceItem_StorageServices interface {
	Save(ii InvoiceItem) (err error)

	Count(invoiceID [16]byte) (numbers uint64, err error)
	Get(invoiceID [16]byte, versionOffset uint64) (ii InvoiceItem, err error)

	// FindByProductID(quiddityID uint64, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
