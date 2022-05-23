

package protocol

import (
	"../libgo/protocol"
)

type FinancialCredit interface {
	UserID() [16]byte                // user-status domain
	Amount() protocol.AmountOfMoney  //
	Balance() protocol.AmountOfMoney //
	RequestID() [16]byte             // user-request domain
}


// do we need this in our software ?

type FinancialCredit_StorageServices interface {
	Save(fc FinancialCredit) (err error)

	Count(UserID [16]byte) (numbers uint64, err error)
	Get(UserID [16]byte, versionOffset uint64) (fc FinancialCredit, err error)
	Last(UserID [16]byte) (fc FinancialCredit, err error)
}