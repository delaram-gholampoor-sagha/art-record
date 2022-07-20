/* For license and copyright information please see LEGAL file in repository */

package art

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
	Get(accountID [16]byte, versionOffset uint64) (fba FinancialBankAccount, nv protocol.NumberOfVersion,err protocol.Error)
	

	FindByIBAN(iban iso.IBAN) (accountID [16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte, nv protocol.NumberOfVersion,err protocol.Error)
}

type (
	FinancialBankAccount_Service_Register_Request interface {
		AccountID() [16]byte
		IBAN() iso.IBAN       // International Bank Account Number
		CardNumber() iso.Card // fixed size card number without any dash or space e.g. 1234123412341234
		Currency() uint64
	}
	
	FinancialBankAccount_Service_Register_Response = protocol.NumberOfVersion

)



type (
	FinancialBankAccount_Service_Get_Request interface {
		AccountID() [16]byte
		versionOffset() uint64
	}

	FinancialBankAccount_Service_Get_Response1 = 	FinancialBankAccount
	FinancialBankAccount_Service_Get_Response2 = protocol.NumberOfVersion
	
)

type (
	FinancialBankAccount_Service_Count_Request interface {
		AccountID() [16]byte
	}

	FinancialBankAccount_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	FinancialBankAccount_Service_FindByIBAN_Request interface {
		Iban() iso.IBAN
	}
	FinancialBankAccount_Service_FindByIBAN_Response1 interface {
		AccountID() [16]byte
	}

	FinancialBankAccount_Service_FindByIBAN_Response2 = protocol.NumberOfVersion
)

type (
	FinancialBankAccount_Service_FindByCardNumber_Request interface {
		CardNumber() iso.Card
	
	}
	FinancialBankAccount_Service_FindByCardNumber_Response1 interface {
		AccountID() [16]byte
	}

	FinancialBankAccount_Service_FindByCardNumber_Response2 = protocol.NumberOfVersion

)