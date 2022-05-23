package protocol

import "time"



type UserOtherAppToken interface {
	UserOtherAppID() [16]byte // user-other-app domain
	Token() string            // DelegateToken , AccessToken
	Time() time.Time          // Save time
	RequestID() [16]byte      // user-request domain
}



type UserOtherAppToken_StorageServices interface {
	Save(uoat UserOtherAppToken) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte, versionOffset uint64) (uoat UserOtherAppToken, err error)
	Last(userID [16]byte) (uoat UserOtherAppToken, length uint64, err error)
} 