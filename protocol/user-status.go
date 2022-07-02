package protocol


type UserStatus interface {
	UserID() [16]byte    // Generated unique ID for the user
	Status() User_Status //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type UserStatus_StorageServices interface {
	Save(us UserStatus) (numbers uint64, err protocol.Error)

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (us UserStatus, numbers uint64, err protocol.Error)

	GetIDs(offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)

	// FindByTime(day utc.DayElapsed, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)

	// protocol.EventTarget
}

type User_Status Quiddity_Status

const (
// User_Status_ = User_Status(Quiddity_Status_FreeFlag << iota)
)
