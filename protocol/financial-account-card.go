package protocol


type FinancialAccountCard interface {
	AccountID() [16]byte       // financial-account domain
	CardNumber() iso.Card      //
	ExpireDate() protocol.Time //
	CVC() uint16               //
	Time() protocol.Time       // Save time
	RequestID() [16]byte       // user-request domain
}

type FinancialAccountCard_StorageServices interface {
	Save(fai FinancialAccountCard) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountCard, err protocol.Error)
	Last(accountID [16]byte) (fai FinancialAccountCard, numbers uint64, err protocol.Error)

	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte, err protocol.Error)
}

type FinancialAccountCard_Service_Register_Request interface {
	AccountID() [16]byte
	CardNumber() iso.Card
	ExpireDate() protocol.Time
	CVC() uint16
}

type FinancialAccountCard_Service_GetLast_Request interface {
	AccountID() [16]byte
}
type FinancialAccountCard_Service_GetLast_Response interface {
	FinancialAccountCard
}
