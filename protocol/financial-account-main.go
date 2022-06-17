package protocol

type FinancialAccountMain interface {
	UserID() [16]byte    // user domain
	Currency() [16]byte  // financial-currency
	AccountID() [16]byte // financial-account domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type FinancialAccountMain_StorageServices interface {
	Save(fa FinancialAccountMain) (err protocol.Error)

	Count(userID, currency [16]byte) (numbers uint64, err protocol.Error)
	Get(userID, currency [16]byte, versionOffset uint64) (fa FinancialAccountMain, err protocol.Error)
	Last(userID, currency [16]byte) (fa FinancialAccountMain, numbers uint64, err protocol.Error)
}
