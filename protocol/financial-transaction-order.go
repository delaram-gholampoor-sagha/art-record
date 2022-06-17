package protocol


// FinancialTransactionOrder indicate the domain record data fields.
type FinancialTransactionOrder interface {
	ID() [16]byte          // FinancialTransactionOrderID
	UserID() [16]byte      // user domain
	ReferenceID() [16]byte // Other Platform transaction ID e.g. bank
	PlatformID() [4]byte   // Detect to||from Currency
	Amount() float64       //
	SuggestPrice() float64 //
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}
