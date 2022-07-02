package protocol


// InvoiceItemSeat indicate the domain record data fields.
type InvoiceItemSeat interface {
	InvoiceID() [16]byte            // invoice domain
	ProductID() [16]byte            // product domain
	SeatID() uint64                 // seat number
	Status() InvoiceItemSeat_Status //
	Time() protocol.Time            // save time
	RequestID() [16]byte            // user-request domain
}

type InvoiceItemSeat_StorageServices interface {
	Save(iis InvoiceItemSeat) (numbers uint64, err protocol.Error)

	Count(invoiceID [16]byte) (numbers uint64, err protocol.Error)
	Get(invoiceID [16]byte, versionOffset uint64) (iis InvoiceItemSeat, numbers uint64, err protocol.Error)
}

type InvoiceItemSeat_Status Quiddity_Status

const (
	InvoiceItemSeat_Status_Voided = InvoiceItemSeat_Status(Quiddity_Status_FreeFlag << iota)
)

// service: reserve seat
