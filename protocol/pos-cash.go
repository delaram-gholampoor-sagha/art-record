package protocol

// PosProvider or pos operator is to supply a particular service.
type PosCash interface {
	PosID() [16]byte     // pos-status domain
	AccountID() [16]byte // financial-account domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
