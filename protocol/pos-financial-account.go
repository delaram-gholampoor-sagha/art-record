package protocol

// PosFinancialAccount or pos operator is to supply a particular service.
type PosFinancialAccount interface {
	PosID() [16]byte     // pos domain
	AccountID() [16]byte // financial-account domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
