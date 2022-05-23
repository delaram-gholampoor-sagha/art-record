

package protocol

/*
	Storage Data Structure
*/

// UserBlock indicate the user-block domain record data fields.
type UserBlock interface {
	UserID() [16]byte        // user-status domain
	BlockUserID() [16]byte   // user-status domain
	Status() UserBlockStatus //
	RequestID() [16]byte     // user-request domain
}


type UserServices_StorageServices interface {
	Save(ub UserBlock) error

	Count(userID [16]byte, BlockUserID [16]byte) (length uint64, err error)
	Get(userID [16]byte, BlockUserID [16]byte, versionOffset uint64) (ub UserBlock, err error)
	Last(userID [16]byte, BlockUserID [16]byte) (ub UserBlock, err error)

	FindUserBlockByStatus(userID [16]byte, BlockUserID [16]byte) (status uint8 , err error)
}

// UserBlockStatus indicate UserBlock record status
type UserBlockStatus uint8

// User status
const (
	UserBlockStatus_Unset UserBlockStatus = iota
	UserBlockStatus_Blocked
	UserBlockStatus_Unblocked
)
