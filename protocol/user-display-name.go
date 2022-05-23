

package protocol

// UserDisplayName indicate the user-display-name domain record data fields.
// It is not unique and not replace of user ID. Just use to find user by desire name.
type UserDisplayName interface {
	UserID() [16]byte    // user-status domain
	DisplayName() string //
	RequestID() [16]byte // user-request domain
}


type UserDisplayName_StorageServices interface {
	Save(udn UserDisplayName) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte,  versionOffset uint64) (udn UserDisplayName, err error)
	Last(userID [16]byte ) (udn UserDisplayName, err error)
}