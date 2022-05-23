package protocol

import (
	"time"

	"../libgo/protocol"
)

// UserRole indicate the user-role domain record data fields.
type UserRole interface {
	ID() [16]byte                          //
	UserID() [16]byte                      // user-status domain
	QuiddityID() [16]byte                  // quiddity domain. Almost use to show locale title.
	GroupID() [16]byte                     // group domain
	AccessControl() protocol.AccessControl //
	Status() UserRole_Status               //
	Time() time.Time                   // Save time
	RequestID() [16]byte                   // user-request domain
}

type UserRole_StorageServices interface {
	Save(ur UserRole) error

	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (ur UserRole, err error)
	Last(id [16]byte) (ur UserRole, err error)

	FindByUserID(userID [16]byte, offset, limit uint64) (ids [][16]byte, length uint64, err error)
	FindByQuiddityID(quiddityID [16]byte) (id [16]byte, err error)
	FindByGroupID(groupID [16]byte, offset, limit uint64) (ids [][16]byte, length uint64, err error)
}

type UserRole_Status uint8

const (
	UserRole_Status_Unset UserRole_Status = iota
	UserRole_Status_Active
	UserRole_Status_Inactive
	UserRole_Status_Deleted
)
