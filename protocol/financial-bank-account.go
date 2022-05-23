

package protocol

import (
	"../libgo/iso"
	"../libgo/protocol"
)

type BankAccount interface {
	AccountID() [16]            // 
	UserID() [16]byte           // user-status domain
	IBAN() iso.IBAN             // International Bank Account Number
	CardNumber() iso.Card       // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64           //
	Status() BankAccount_Status //
	RequestID() [16]byte        // user-request domain
}

type BankAccount_StorageServices interface {
	Save(ba BankAccount) (err error)
r
	Count(accountID [16]byte) (numbers uint64, err error)
	Get(accountID [16]byte, versionOffset uint64) (ba BankAccount, err error)
	Last(accountID [16]byte) (ba BankAccount, err error)

	FindByIBAN(iban string) (accountID [16]byte, err error)
	FindByCardNumber(cardNumber string) (accountID [16]byte, err error)

	ListUserAccounts(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err error)
}

// BankAccount_Status indicate BankAccount record status
type BankAccount_Status uint8

// BankAccount status
const (
	BankAccount_Status_Unset BankAccountStatus = 0b00000000
	BankAccount_Status_Block BankAccountStatus = 0b00000001
	BankAccount_Status_Close BankAccountStatus = 0b00000010
)
