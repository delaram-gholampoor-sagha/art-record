/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type UserStatus interface {
	UserID() [16]byte    // Generated unique ID for the user
	Status() User_Status //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type UserStatus_StorageServices interface {
	Save(us UserStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(userID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (us UserStatus, nv protocol.NumberOfVersion, err protocol.Error)

	GetIDs(offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	// FindByTime(day utc.DayElapsed, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	// protocol.EventTarget
}

type User_Status Quiddity_Status

const (
// User_Status_ = User_Status(Quiddity_Status_FreeFlag << iota)
)



type (
	UserStatus_Service_Register_Request interface {
		Status() User_Status
	}

	UserStatus_Service_Register_Response1 interface {
		UserID() [16]byte
	}

	UserStatus_Service_Register_Response2 = protocol.NumberOfVersion
)


type (

	UserStatus_Service_Count_Request interface {
		UserID() [16]byte
	}

	UserStatus_Service_Count_Response = protocol.NumberOfVersion
)


type (
	UserStatus_Service_Get_Request interface {
		UserID() [16]byte
		VersionOffset() uint64
	}

	UserStatus_Service_Get_Response1 = UserStatus
	UserStatus_Service_Get_Response2 = protocol.NumberOfVersion
)


type (
	UserStatus_Service_GetIDs_Request interface {
		Offset() uint64
		Limit() uint64
	}

	UserStatus_Service_GetIDs_Response1 interface {
		UserIDs() [][16]byte
	}

	UserStatus_Service_GetIDs_Response2 = protocol.NumberOfVersion
)