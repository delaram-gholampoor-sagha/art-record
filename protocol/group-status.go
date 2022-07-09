package protocol

import "../libgo/protocol"

// GroupStatus indicate the domain record data fields.
type GroupStatus interface {
	GroupID() [16]byte    // group-status domain
	Status() Group_Status //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type GroupStatus_StorageServices interface {
	Save(gs GroupStatus) (numbers uint64, err protocol.Error)

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gs GroupStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status Picture_Status, offset, limit uint64) (groupIDs [][16]byte, numbers uint64, err protocol.Error)

	// protocol.EventTarget
}

type (
	GroupStatus_Service_Register_Request interface {
		GroupID() [16]byte    
		Status() Group_Status 
	}
	GroupStatus_Service_Register_Response interface {
		Numbers() uint64
	
	}
	
	GroupStatus_Service_Count_Request interface { 
		GroupID() [16]byte 
	
	}
	GroupStatus_Service_Count_Response interface {
		Numbers() uint64
	}
	GroupStatus_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	
	
	}
	GroupStatus_Service_Get_Response interface {
		GroupStatus
		Numbers() uint64
	}
)

type Group_Status Quiddity_Status

const (
// Group_Status_ = Group_Status(Quiddity_Status_FreeFlag << iota)
)
