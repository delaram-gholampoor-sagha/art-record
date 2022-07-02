package protocol

// UserReferent indicate the domain record data fields.
type UserReferent interface {
	UserID() [16]byte         // user domain
	ReferentUserID() [16]byte // user domain
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type UserReferent_StorageServices interface {
	Save(ur UserReferent) (numbers uint64, err protocol.Error)

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (ur UserReferent, numbers uint64, err protocol.Error)

	FindByReferentUserID(referentUserID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
}
