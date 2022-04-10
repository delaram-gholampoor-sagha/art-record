package protocol





type FinancialTransaction interface {
	AccountID() [16]byte   
	// communicator
	SenderAccountID() [16]byte      
	Reference() [16]byte            // Any ID that user can assign to track for any purpose
	Amount() int64                  // Some number base on currency is Decimal part e.g. 8099 >> 80.99$
	Balance() int64                 // Some number base on currency is Decimal part e.g. 8099 >> 80.99$
	Type() FinancialTransactionType //
	RequestID() [16]byte            // Request ObjectID to prove how this object created in the chain of other objects.
}



type FinancialTransactionServices interface {
	LastTransaction(accountID [16]byte) (ba FinancialTransaction, err  error)
	LastTransactionVersion(accountID [16]byte) (versionOffset uint64, version int64, err error)
	TransactionVersions(accountID [16]byte, day protocol.TimeUnixDay, offset, limit uint64) (ver []int64, err error)
	GetTransaction(accountID [16]byte, day protocol.TimeUnixDay, versionOffset uint64) (ba FinancialTransaction, err error)

	Save(fa FinancialTransaction) (version int64, err error)
}

/*
	Enumerated types
*/

// FinancialTransactionType indicate FinancialTransaction record type
type FinancialTransactionType uint8

// FinancialTransaction types
const (
	FinancialTransactionUnset  FinancialTransactionType = iota
	FinancialTransactionFailed                          // For any reason e.g. not same currency account, not reachable bank, ...
	FinancialTransactionBlocked
	FinancialTransactionExchange
	FinancialTransactionPrize
	FinancialTransactionDonate
	FinancialTransactionCharity
	FinancialTransactionTip
	FinancialTransactionCash          // >> Due to digital era banking, prefer to not support cash transactions
	FinancialTransactionRevolvingFund // >> Due to digital era banking, prefer to not support this type
	// FinancialTransactionSameBankTransfer     >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransactionForeignBankTransfer  >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransactionPOSTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	// FinancialTransactionWebTransfer          >> Not the goal and just the way of transaction that can check by RequestID
	FinancialTransactionAsset
	FinancialTransactionStock
	FinancialTransactionRentalFee
	FinancialTransactionBill
	FinancialTransactionInvoice
	FinancialTransactionReturnInvoice
	FinancialTransactionTax
	FinancialTransactionCommission
	FinancialTransactionSalary // Reward
	FinancialTransactionProfit
	FinancialTransactionLoan
	FinancialTransactionBadDebt // Debt default
	FinancialTransactionInstallment
	FinancialTransactionCompensation
)

/*
	Services
*/

const FinancialTransactionServiceRegister = "urn:giti:financial-transaction.protocol:service:transfer"

const FinancialTransactionServiceGetLast = "urn:giti:financial-transaction.protocol:service:get-last"

type GetLastFinancialTransactionRequest interface {
	RequestID() [16]byte            // Request ObjectID to prove how this object created in the chain of other objects.
	AccountID() [16]byte            //
	SenderAccountID() [16]byte      //
	Reference() [16]byte            // Any ID that user can assign to track for any purpose
	Amount() protocol.Amount        // Some number base on currency is Decimal part e.g. 8099 >> 80.99$
	Balance() protocol.Amount       // Some number base on currency is Decimal part e.g. 8099 >> 80.99$
	Type() FinancialTransactionType //
}

const FinancialTransactionServiceGetByTime = "urn:giti:financial-transaction.protocol:service:get-by-time"
const FinancialTransactionServiceFind = "urn:giti:financial-transaction.protocol:service:find"

/*
	Errors
*/
