

package protocol

// UserReferent indicate the user-referent domain record data fields.
type UserReferent interface {
	UserID() [16]byte         // user-status domain
	ReferentUserID() [16]byte // user-status domain
	RequestID() [16]byte      // user-request domain
}


type UserReferent_StorageServices interface {
	Save(ur UserReferent) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte, versionOffset uint64) (ur UserReferent, err error)
	Last(userID [16]byte) (ur UserReferent, length uint64, err error)

	
}
