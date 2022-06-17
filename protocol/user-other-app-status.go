package protocol

type UserOtherAppStatus interface {
	UserOtherAppID() [16]byte    // user-other-app domain
	Status() UserOtherApp_Status //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type UserOtherAppStatus_StorageServices interface {
	Save(uos UserOtherAppStatus) protocol.Error

	Count(userOtherAppID [16]byte) (numbers uint64, err protocol.Error)
	Get(userOtherAppID [16]byte, versionOffset uint64) (uos UserOtherAppStatus, err protocol.Error)
	Last(userOtherAppID [16]byte) (uos UserOtherAppStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status UserOtherApp_Status, offset, limit uint64) (userOtherAppIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type UserOtherApp_Status Quiddity_Status

const (
// UserOtherApp_Status_ UserOtherApp_Status = (Quiddity_Status_FreeFlag << iota)
)
