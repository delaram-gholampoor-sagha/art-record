package protocol

import (
	"../libgo/protocol"
)

// RoleStaff indicate the domain record data fields.
type RoleStaff interface {
	RoleID() [16]byte    // role domain
	Minimum() int        // Minimum number of employees for this job position
	Expected() int       // Expected number of employees for this job position
	Maximum() int        // Maximum number of employees for this job position
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type RoleStaff_StorageServices interface {
	Save(rs RoleStaff) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rs RoleStaff, err protocol.Error)
	Last(roleID [16]byte) (rs RoleStaff, numbers uint64, err protocol.Error)
}
