package protocol

import "time"

type UserOtherApp interface {
	ID() [16]byte         //
	UserID() [16]byte     // user-status domain
	Domain() string       // unique name in the computer world. Country e.g. Iran, ... || Digital platforms e.g. Instagram, Facebook, Google, Github, Yahoo, ...
	DomainUserID() string // user id in that domain e.g. country national code
	Time() time.Time      // Save time
	RequestID() [16]byte  // user-request domain
}


type UserOtherApp_StorageServices interface {
	Save(uoa UserOtherApp) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte, versionOffset uint64) (uoa UserOtherApp, err error)
	Last(userID [16]byte) (uoa UserOtherApp, length uint64, err error)
}


