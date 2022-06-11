package protocol

import (
	"../libgo/protocol"
)

type UserOtherAppStatus interface {
	UserOtherAppID() [16]byte    // user-other-app domain
	Status() UserOtherApp_Status //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

type UserOtherApp_Status uint8

const (
	UserOtherApp_Status_Unset UserOtherApp_Status = iota
	UserOtherApp_Status_Registered
	UserOtherApp_Status_Inactivated
	UserOtherApp_Status_Blocked
)
