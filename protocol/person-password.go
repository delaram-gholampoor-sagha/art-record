package protocol

import "time"

// PersonPassword indicate the domain record data fields.
// A person can decide to don't register password and deny authenticate by password.
type PersonPassword interface {
	PersonID() [16]byte            // user domain
	PasswordHash() []byte          //
	Status() PersonPassword_Status //
	RequestID() [16]byte           // user-request domain
}

type PersonPassword_StorageServices interface {
	Create(password Password) error
	Last(userID [16]byte) (Password, error)
	Lasts(userID [16]byte, count int) ([]Password, error)
	Meantime(userID [16]byte, start, end time.Time) ([]Password, error)
}

// PersonPassword_Status indicate PersonPassword record status
type PersonPassword_Status uint8

// PersonPassword_Status status
const (
	PersonPassword_Status_Unset PersonPassword_Status = 0
	// authenticate person with Password + OTP
	PersonPassword_Status_ForceUse2Factor PersonPassword_Status = (1 << iota)
	// user must change password and otp key(if enabled) for some reason
	PersonPassword_Status_MustChangePassword
	// user can't use or decide to prevent authenticate via password with or without OTP.
	PersonPassword_Status_DenyPasswordAuthentication
)

type PersonPassword_Service_SendResetToken_Request interface {
	How() uint8 // sms, email, ...
}
type PersonPassword_Service_SendResetToken_Response interface {
	Token() []byte // include Token&ExpireIn with signature by domain key.
}

type PersonPassword_Service_Register_Request interface {
	Password() []byte
}

type PersonPassword_Service_Get_Request interface {
	VersionOffset() uint64
}
type PersonPassword_Service_Get_Response interface {
	PersonPassword
}

type PersonPassword_Service_GetLast_Request interface {
	PersonPassword
	Numbers() uint64
}
