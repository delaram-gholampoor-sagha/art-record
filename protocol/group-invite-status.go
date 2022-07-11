package protocol

import "../libgo/protocol"

// GroupInviteStatus indicate the domain record data fields.
type GroupInviteStatus interface {
	GroupID() [16]byte          // group domain
	UserID() [16]byte           // user domain
	Status() GroupInvite_Status //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}


type GroupInviteStatus_StorageServices interface {
	Save(gn GroupInviteStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gn GroupInviteStatus, nv protocol.NumberOfVersion, err protocol.Error)

	
}

type GroupInvite_Status uint8

const (
	GroupInvite_Status_Unset GroupInvite_Status = iota
	GroupInvite_Status_Active
	GroupInvite_Status_Inactive
)

// TODO::: user send last time of active state record plus its ID and optional invited user id as invite code.

type (
	
	GroupInviteStatus_Service_Register_Request interface {
		GroupID() [16]byte          
		UserID() [16]byte          
		Status() GroupInvite_Status 
	}
	GroupInviteStatus_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	GroupInviteStatus_Service_Count_Request interface { 
		GroupID() [16]byte
	}
	GroupInviteStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	GroupInviteStatus_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	}
	GroupInviteStatus_Service_Get_Response interface {
		GroupInviteStatus
		Nv() protocol.NumberOfVersion
	}

)