/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

// Role indicate the role-status domain record data fields.
type RoleStatus interface {
	RoleID() [16]byte    // role domain
	Status() Role_Status //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type RoleStatus_StorageServices interface {
	Save(rs RoleStatus) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rs RoleStatus, err protocol.Error)
	Last(roleID [16]byte) (rs RoleStatus, err protocol.Error)
}

type Role_Status uint8

const (
	Role_Status_Unset Role_Status = iota
	Role_Status_Active
	Role_Status_Inactive
	Role_Status_Deleted
	Role_Status_Revoked

	Role_Status_JobOpening
)
