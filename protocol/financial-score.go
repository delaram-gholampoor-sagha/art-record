package protocol

// Individual and General score
type FinancialScore interface {
	UserID() [16]byte    // user domain
	Score() int64        //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
