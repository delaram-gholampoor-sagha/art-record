package protocol

import (
	"../libgo/iso"
	"../libgo/protocol"
)

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

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fba FinancialBankAccount, err protocol.Error)
	Last(accountID [16]byte) (fba FinancialBankAccount, numbers uint64, err protocol.Error)

	FindByIBAN(iban iso.IBAN) (accountID [16]byte, err protocol.Error)
	FindByCardNumber(cardNumber iso.Card) (accountID [16]byte, err protocol.Error)
}

type FinancialBankAccount_Service_Register_Request interface {
	AccountID() [16]byte
	IBAN() iso.IBAN       // International Bank Account Number
	CardNumber() iso.Card // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64
}

type FinancialBankAccount_Service_GetLast_Request interface {
	FinancialBankAccount
}
