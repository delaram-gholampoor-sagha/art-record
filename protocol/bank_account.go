package protocol




type BankAccount interface {
	AccountID() [16]byte
	UserID() [16]byte   //
	IBAN() string       // International Bank Account Number
	CardNumber() string // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64
	Status() BankAccountStatus
	RequestID() [16]byte // Request ObjectID to prove how this object created in the chain of other objects.
}



type BankAccountServices interface {
	LastAccountDetail(accountID [16]byte) (ba BankAccount, err error)
	LastAccountVersion(accountID [16]byte) (versionOffset uint64, version int64, err error)
	AccountVersions(accountID [16]byte, offset, limit uint64) (ver []int64, err error)
	GetAccount(accountID [16]byte, versionOffset uint64) (ba BankAccount, err error)

	FindAccountByIBAN(iban string) (accountID [16]byte, err error)
	FindAccountByCardNumber(cardNumber string) (accountID [16]byte, err error)

	ListUserAccounts(userID [16]byte, offset, limit uint64) (accountID [][16]byte, err error)

	Save(fa BankAccount) (version int64, err error)
}

/*
	Enumerated types
*/

// BankAccountStatus indicate BankAccount record status
type BankAccountStatus uint8

// BankAccount status
const (
	BankAccountStatusUnset BankAccountStatus = 0b00000000
	BankAccountStatusBlock BankAccountStatus = 0b00000001
	BankAccountStatusClose BankAccountStatus = 0b00000010
)

/*
	Services
*/

const (
// ApplicationService = "urn:giti:application.protocol:service:"
)

const BankAccountServiceGetLast = "urn:giti:financial-transaction.protocol:service:get-last"

type RegisterNewBankAccountRequest interface {
	IBAN()       // International Bank Account Number
	CardNumber()  // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64
}

type LastBankAccountDetailResponse interface {
	RequestID() [16]byte // Request ObjectID to prove how this object created in the chain of other objects.
	AccountID() [16]byte
	UserID() [16]byte     //
	IBAN()        // International Bank Account Number
	CardNumber() // fixed size card number without any dash or space e.g. 1234123412341234
	Currency() uint64
	Status() BankAccountStatus
}

/*
	Errors
*/

const (
// ApplicationError = "urn:giti:application.protocol:error:"
)