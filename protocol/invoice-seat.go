package protocol

type InvoiceSeat interface {
	InvoiceID() [16]byte // invoice-status domain
	ProductID() [16]byte // product domain
	SeatID() uint64      // seat number in the specific raw
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

// reserve seat
