package protocol

type FinancialCurrency interface {
	Currency() [16]byte // quiddity domain
	// ISO_4217() iso.Currency             //
	MoneySettlementReference() [16]byte // user-status domain. can be any org or society user.
	Time() protocol.Time                // Save time
	RequestID() [16]byte                // user-request domain
}
