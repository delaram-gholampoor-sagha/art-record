package protocol

import (
	"../libgo/protocol"
)

// UserDisplayName indicate the domain record data fields.
// It is not unique and not replace of user ID. Just use to find user by desire name.
type UserDisplayName interface {
	UserID() [16]byte    // user-status domain
	DisplayName() string //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
