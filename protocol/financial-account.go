package protocol


type FinancialAccount interface {
	AccountID() [16]byte                // quiddity domain
	Currency() [16]byte                 // financial-currency
	UserID() [16]byte                   // user-status domain
	MoneySettlementReference() [16]byte // user-status domain. can be any org or society user.
	Time() protocol.Time                // Save time
	RequestID() [16]byte                // user-request domain
}

type FinancialAccount_StorageServices interface {
	Save(fa FinancialAccount) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fa FinancialAccount, err protocol.Error)
	Last(accountID [16]byte) (fa FinancialAccount, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (accountIDs [][16]byte, numbers uint64, err protocol.Error)

	ListSettlementReferences(userID [16]byte, offset, limit uint64) (moneySettlementReferences [][16]byte, numbers uint64, err protocol.Error)
}
