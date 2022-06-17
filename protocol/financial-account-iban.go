package protocol
type FinancialAccountIBAN interface {
	AccountID() [16]byte // financial-account domain
	IBAN() iso.IBAN      //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type FinancialAccountIBAN_StorageServices interface {
	Save(fai FinancialAccountIBAN) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fai FinancialAccountIBAN, err protocol.Error)
	Last(accountID [16]byte) (fai FinancialAccountIBAN, numbers uint64, err protocol.Error)

	FindByIBAN(iban iso.IBAN) (accountID [16]byte, err protocol.Error)
}

type FinancialAccountIBAN_Service_Register_Request interface {
	AccountID() [16]byte
	IBAN() iso.IBAN
}

type FinancialAccountIBAN_Service_GetLast_Request interface {
	AccountID() [16]byte
}
type FinancialAccountIBAN_Service_GetLast_Response interface {
	FinancialAccountIBAN
}
