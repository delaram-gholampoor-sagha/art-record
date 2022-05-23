package protocol

/*
	Storage Data Structure
*/

// UserPublicKey indicate the user-public-key domain record data fields.
type UserPublicKey interface {
	UserID() [16]byte            // user-status domain
	PublicKey() []byte           // suggest use DER format
	Status() UserPublicKeyStatus //
	RequestID() [16]byte         // user-request domain
}

type UserPublicKey_StorageServices interface {
	Save(up UserPublicKey) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte, versionOffset uint64) (up UserPublicKey, err error)
	Last(userID [16]byte) (up UserPublicKey, length uint64, err error)
}



// UserPublicKeyStatus use to indicate UserPublicKey record status
type UserPublicKeyStatus uint8

// UserPublicKey status
const (
	UserPublicKeyStatus_Unset UserPublicKeyStatus = iota
	UserPublicKeyStatus_Active
	UserPublicKeyStatus_Inactive
	UserPublicKeyStatus_Blocked
	UserPublicKeyStatus_Revoked
)
