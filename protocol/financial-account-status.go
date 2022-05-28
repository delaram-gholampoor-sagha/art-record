package protocol

import (
	"../libgo/protocol"
)

type FinancialAccountStatus interface {
	AccountID() [16]byte             //
	Status() FinancialAccount_Status //
	Time() protocol.Time             // Save time
	RequestID() [16]byte             // user-request domain
}

type FinancialAccountStatus_StorageServices interface {
	Save(fas FinancialAccountStatus) (err protocol.Error)

	Count(accountID [16]byte) (numbers uint64, err protocol.Error)
	Get(accountID [16]byte, versionOffset uint64) (fas FinancialAccountStatus, err protocol.Error)
	Last(accountID [16]byte) (fas FinancialAccountStatus, numbers uint64, err protocol.Error)
}

type FinancialAccount_Status uint8

const (
	FinancialAccount_Status_Unset      FinancialAccount_Status = 0
	FinancialAccount_Status_Registered FinancialAccount_Status = (1 << iota)
	FinancialAccount_Status_Closed
	FinancialAccount_Status_Blocked
)
