package protocol

import "../libgo/protocol"

type FinancialBankAccount interface {
	AccountID() [16]byte  // financial-account-status domain
	IBAN() iso.IBAN       // International Bank Account Number
	CardNumber() iso.Card // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64     //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type FinancialBankAccount_StorageServices interface {
	Save(fba FinancialBankAccount) (err protocol.Error)

	Count(accountID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fba FinancialBankAccount, err protocol.Error)
	

	FindByIBAN(iban iso.IBAN) (accountID [16]byte, err protocol.Error)
	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte, err protocol.Error)
}

type (
	FinancialBankAccount_Service_Register_Request interface {
		AccountID() [16]byte
		IBAN() iso.IBAN       // International Bank Account Number
		CardNumber() iso.Card // fixed size card number without any dash or space e.g. 1234123412341234
		Currency() uint64
	}
	
	FinancialBankAccount_Service_Register_Response interface {
		Numbers() uint32
	}
	
	
	FinancialBankAccount_Service_GetLast_Request interface {
		AccountID() [16]byte
	}
	FinancialBankAccount_Service_GetLast_Response interface {
		FinancialBankAccount
	}
	
	
	FinancialBankAccount_Service_Get_Request interface {
		AccountID() [16]byte
		VersionOffset() uint64
	}
	FinancialBankAccount_Service_Get_Response interface {
		FinancialBankAccount
	}
	
	
	FinancialBankAccount_Service_Count_Request interface {
		AccountID() [16]byte
	
	}
	FinancialBankAccount_Service_Count_Response interface {
		Numbers() uint32
	}
	
	FinancialBankAccount_Service_FindByIBAN_Request interface {
		Iban() iso.IBAN
	
	}
	FinancialBankAccount_Service_FindByIBAN_Response interface {
		AccountID() [16]byte
	}
	
	FinancialBankAccount_Service_FindByCardNumber_Request interface {
		CardNumber() iso.Card
	
	}
	FinancialBankAccount_Service_FindByCardNumber_Response interface {
		AccountID() [16]byte
	}
)