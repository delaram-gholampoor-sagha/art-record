
package protocol

import (
	"../libgo/protocol"
)

// UserDelegate indicate the user-delegate domain record data fields.
type UserDelegate interface {
	UserID() [16]byte            // user-status domain
	DelegateUserID() [16]byte    // user-status domain
	UserRoleID() [16]byte        // user-role domain
	Status() UserDelegate_Status //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

type UserDelegate_StorageServices interface {
	Save(ud UserDelegate) error

	Count(userID [16]byte, delegateUserID [16]byte) (numbers uint64, err error)
	Get(userID [16]byte, delegateUserID [16]byte, versionOffset uint64) (ud UserDelegate, err error)
	Last(userID [16]byte, delegateUserID [16]byte) (ud UserDelegate, err error)

	FindByDelegateUserID(delegateUserID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err error)
	FindByRole(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err error)

	ListGivenDelegate(userID [16]byte, offset, limit uint64) (delegateUserIDs [][16]byte, numbers uint64, err error)
}

type UserDelegate_Status uint8

const (
	UserDelegate_Status_Unset UserDelegate_Status = iota
	UserDelegate_Status_Deleted
	UserDelegate_Status_Revoked
)
