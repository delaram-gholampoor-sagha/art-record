package protocol

import "time"



type UserOtherAppStatus interface {
	UserOtherAppID() [16]byte          // user-other-app domain
	Status() UserOtherAppStatus_Status //
	Time() time.Time                   // Save time
	RequestID() [16]byte               // user-request domain
}

type UserOtherAppStatus_StorageServices interface {
	Save(uoas UserOtherAppStatus) error

	Count(userID [16]byte) (length uint64, err error)
	Get(userID [16]byte, versionOffset uint64) (uoas UserOtherAppStatus, err error)
	Last(userID [16]byte) (uoas UserOtherAppStatus, length uint64, err error)
} 


type UserOtherAppStatus_Status uint8

const (
	UserOtherAppStatus_Status_Unset UserOtherAppStatus_Status = iota
	UserOtherAppStatus_Status_Registered
	UserOtherAppStatus_Status_Inactivated
	UserOtherAppStatus_Status_Blocked
)
