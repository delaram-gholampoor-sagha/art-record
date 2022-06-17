package protocol

// PosProvider or pos operator is to supply a particular service.
type PosCash interface {
	PosID() [16]byte                    // pos domain
	AccountID() [16]byte                // financial-account domain
	MaxCash() protocol.AmountOfMoney    //
	MaxCashTransfer() protocol.Duration // from reach the max limit
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}
