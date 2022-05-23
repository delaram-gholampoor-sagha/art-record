

package protocol



// Invoice indicate the invoice domain record data fields.
// Also usually known as shopping-cart, ... in GUI part of applications
// Ship service can add as an item to invoice
// each Invoice record is immutable record and we don't use version mechanism.
type Invoice interface {
	ID() [16]byte        // InvoiceID
	CreatorID() [16]byte // user-status domain. Creator of invoice e.g. restaurant garson,
	OwnerID() [16]byte   // user-status domain. Customer user
	RequestID() [16]byte // user-request domain
}

type Invoice_StorageServices interface {
	Save(i Invoice) error
	Update(i Invoice) error  // just use to add OwnerID later than create time

	Get(id [16]byte) (i Invoice, err error)

	GetIDs(offset, limit uint64) (ids [][16]byte, numbers uint64, err error)
	// GetIDsByDateTime(time protocol.Time, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)

	FindByCreatorID(creatorID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err error)
	FindByOwnerID(ownerID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err error)
}
