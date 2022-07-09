package protocol

import "../libgo/protocol"

// GroupOwner indicate the domain record data fields.
type GroupOwner interface {
	GroupID() [16]byte     // or ChannelID. Use to store and retrieve comments in time order. It can be UUID, URL hash or any random id.
	OwnerUserID() [16]byte // user domain
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}

type GroupOwner_StorageServices interface {
	Save(g GroupOwner) (numbers uint64, err protocol.Error)

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (g GroupOwner, numbers uint64, err protocol.Error)

	FindByUserID(ownerUserID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}

type (
	GroupOwner_Service_Register_Request interface {
		GroupID() [16]byte                 
		OwnerUserID() [16]byte    
	}
	GroupOwner_Service_Register_Response interface {
		Numbers() uint64
	}
	
	GroupOwner_Service_Count_Request interface { 
		GroupID() [16]byte
	}
	GroupOwner_Service_Count_Response interface {
		Numbers() uint64
	}
	GroupOwner_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	}
	GroupOwner_Service_Get_Response interface {
		GroupOwner
		Numbers() uint64
	}
	GroupOwner_Service_FindByUserID_Request interface { 
		OwnerUserID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	GroupOwner_Service_FindByUserID_Response interface {
		IDs() [][16]byte
		Numbers() uint64
	}
)