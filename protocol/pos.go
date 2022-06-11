package protocol

// POS or point-of-sale
type POS interface {
	PosID() [16]byte // quiddity domain

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
