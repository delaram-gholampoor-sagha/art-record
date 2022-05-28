package protocol

import (
	"../libgo/protocol"
)

// UserReferent indicate the domain record data fields.
type UserReferent interface {
	UserID() [16]byte         // user-status domain
	ReferentUserID() [16]byte // user-status domain
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
