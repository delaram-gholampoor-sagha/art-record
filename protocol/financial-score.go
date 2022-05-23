

package protocol



type FinancialScore interface {
	UserID() [16]byte    // user-status domain
	Score() int32        //
	RequestID() [16]byte // user-request domain
}


// do we need this in our software ?

type FinancialScore_StorageServices interface {
	Save(fs FinancialScore) (err error)

	Count(UserID [16]byte) (numbers uint64, err error)
	Get(UserID [16]byte, versionOffset uint64) (fs FinancialScore, err error)
	Last(UserID [16]byte) (fs FinancialScore, err error)
}