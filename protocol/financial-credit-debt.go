

package protocol

import (
	"../libgo/protocol"
)

// https://en.wikipedia.org/wiki/Credit_card_debt
type FinancialCreditDebt interface {
	UserID() [16]byte                // user-status domain
	Amount() protocol.AmountOfMoney  //
	Balance() protocol.AmountOfMoney //
	RequestID() [16]byte             // user-request domain
}

// do we need this in our software ?

type FinancialCreditDebt_StorageServices interface {
	Save(fc FinancialCreditDebt) (err error)

	Count(UserID [16]byte) (numbers uint64, err error)
	Get(UserID [16]byte, versionOffset uint64) (fc FinancialCreditDebt, err error)
	Last(UserID [16]byte) (fc FinancialCreditDebt, err error)
}