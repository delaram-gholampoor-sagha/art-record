package protocol

// FinancialCredit indicate the domain record data fields.
type FinancialCredit interface {
	UserID() [16]byte                // user domain
	Currency() [16]byte              // financial-currency
	Amount() protocol.AmountOfMoney  //
	Balance() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}
