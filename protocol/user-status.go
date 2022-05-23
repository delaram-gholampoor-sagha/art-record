

package protocol

/*
	Storage Data Structure
*/

type UserStatus interface {
	UserID() [16]byte    // Generated unique ID for the user
	Status() UserStatus_ //
	RequestID() [16]byte // user-request domain
}


type UserStatus_StorageServices interface {
	Save(us UserStatus) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte, versionOffset uint64) (us UserStatus, err error)
	Last(userID [16]byte) (us UserStatus, length uint64, err error)

}

// UserStatus_ indicate UserStatus record status
type UserStatus_ uint8

// User status
const (
	UserStatus_Unset UserStatus_ = iota
	UserStatus_Active
	UserStatus_Inactive
	UserStatus_Blocked
)
