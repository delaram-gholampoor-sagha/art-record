package protocol
import (
	"../libgo/protocol"
)

type FinancialAccount interface {
	AccountID() [16]byte                // financial-account-status domain
	UserID() [16]byte                   // user-status domain
	SocietyID() [16]byte                // society domain
	MoneySettlementReference() [16]byte // user-status domain
	Time() protocol.Time                // Save time
	RequestID() [16]byte                // user-request domain
}

type FinancialAccount_StorageServices interface {
	Save(fa FinancialAccount) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fa FinancialAccount, err protocol.Error)
	Last(accountID [16]byte) (fa FinancialAccount, numbers uint64, err protocol.Error)

	FindByUserID(userID [16]byte, offset, limit uint64) (accountIDs [][16]byte, numbers uint64, err protocol.Error)

	ListUserSocieties(userID [16]byte, offset, limit uint64) (societyIDs [][16]byte, numbers uint64, err protocol.Error)
}
