package protocol 

import (
	"../libgo/protocol"
)

// FinancialTransactionOrder to 
type FinancialTransactionOrder struct {
	ID() [16]byte                      // FinancialExchangeOrderID
	UserID() [16]byte                  // user-status domain
	ReferenceID  [16]byte // Other Platform transaction ID e.g. bank
	PlatformID   [4]byte  // Detect to||from Currency
	Amount       float64  //
	SuggestPrice float64
}
