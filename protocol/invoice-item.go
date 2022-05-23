

package protocol
import (
	"../libgo/protocol"
)

// InvoiceItem indicate the domain record data fields.
// each InvoiceItem is immutable record and so use version mechanism to chain InvoiceItem in InvoiceID group.
// Ship service can add as an item to invoice
// Each item must check for reserve validity and make action in end of validity
type InvoiceItem interface {
	InvoiceID() [16]byte // invoice-status domain
	ProductID() [16]byte // product-status domain
	Quantity() uint64    // decimal >> 1.5 >> 1.2 >> 10 >> 0
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type InvoiceItem_StorageServices interface {
	Save(ii InvoiceItem) (err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (ii InvoiceItem, err protocol.Error)

	// FindByProductID(productID uint64, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
