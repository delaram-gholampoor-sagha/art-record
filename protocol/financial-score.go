package protocol

// Individual and General score
type FinancialScore interface {
	UserID() [16]byte    // user-status domain
	Score() int64        //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
