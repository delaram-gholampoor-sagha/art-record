package protocol

// MoneySettlementReference: https://www.federalreserve.gov/paymentsystems/fedfunds_about.htm
type FinancialCurrency interface {
	Currency() [16]byte // quiddity domain
	// ISO_4217() iso.Currency             //
	MoneySettlementReference() [16]byte // user domain. can be any org or society user.
	Time() protocol.Time                // save time
	RequestID() [16]byte                // user-request domain
}
