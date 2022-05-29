package protocol

import (
	"../libgo/protocol"
)

// GroupJoinByInvite indicate the domain record data fields.
type InviteStatus interface {
	GroupID() [16]byte                // group-status domain
	UserID() [16]byte                 // user-status domain
	Status() GroupJoinByInvite_Status //
	Time() protocol.Time              // Save time
	RequestID() [16]byte              // user-request domain
}

type GroupJoinByInvite_Status uint8

const (
	GroupJoinByInvite_Status_Unset GroupJoinByInvite_Status = iota
	GroupJoinByInvite_Status_Active
	GroupJoinByInvite_Status_Inactive
)

// TODO::: user send last time of active state record plus its ID as invite code.
