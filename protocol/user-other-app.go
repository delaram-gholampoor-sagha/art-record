package protocol

import (
	"../libgo/protocol"
)

// UserOtherApp store data about user delegation to this org(app) in other orgs(apps)
type UserOtherApp interface {
	ID() [16]byte         //
	UserID() [16]byte     // user-status domain
	Domain() string       // unique name in the computer world. Country e.g. Iran, ... || Digital platforms e.g. Instagram, Facebook, Google, Github, Yahoo, ...
	DomainUserID() []byte // user id in that domain e.g. country national code
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
