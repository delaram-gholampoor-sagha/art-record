package protocol


// UserDelegate indicate the domain record data fields.
type UserDelegate interface {
	UserID() [16]byte            // user domain
	DelegateUserID() [16]byte    // user domain
	UserRoleID() [16]byte        // user-role domain
	Status() UserDelegate_Status //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type UserDelegate_StorageServices interface {
	Save(ud UserDelegate) protocol.Error

	Count(userID [16]byte, delegateUserID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, delegateUserID [16]byte, versionOffset uint64) (ud UserDelegate, err protocol.Error)
	Last(userID [16]byte, delegateUserID [16]byte) (ud UserDelegate, numbers uint64, err protocol.Error)

	FindByDelegateUser(delegateUserID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
	FindByRole(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)

	ListGivenDelegate(userID [16]byte, offset, limit uint64) (delegateUserIDs [][16]byte, numbers uint64, err protocol.Error)
}

type UserDelegate_Status uint8

const (
	UserDelegate_Status_Unset UserDelegate_Status = iota
	UserDelegate_Status_Deleted
	UserDelegate_Status_Revoked
)

// Service: delegate user delegate its role to other. blocked this service by default.
