package protocol


// https://en.wikipedia.org/wiki/Credit_card_debt
type FinancialCreditDebt interface {
	UserID() [16]byte                // user-status domain
	Amount() protocol.AmountOfMoney  //
	Balance() protocol.AmountOfMoney //
	Time() protocol.Time             // Save time
	RequestID() [16]byte             // user-request domain
}
