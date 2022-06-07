package protocol

import (
	"../libgo/protocol"
)

// RoleSkill indicate the domain record data fields.
type RoleSkill interface {
	RoleID() [16]byte // role domain

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
