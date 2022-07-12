package protocol

import "../libgo/protocol"

// GroupRole indicate the domain record data fields.
type GroupSettings interface {
	GroupID() [16]byte                     // group domain
	ArchiveThreadAfter() protocol.Duration //
	BlockedWords() []string                //
	Join() GroupSettings                   //
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type GroupSettings_StorageServices interface {
	Save(gs GroupSettings) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gs GroupSettings, nv protocol.NumberOfVersion, err protocol.Error)
}

type GroupSettings_Join uint8

const (
	GroupSettings_Join_Unset  GroupSettings_Join = 0
	GroupSettings_Join_Invite GroupSettings_Join = (1 << iota)
	GroupSettings_Join_AddByMembers
)

type (
	GroupSettings_Service_Register_Request interface {
		GroupID() [16]byte                     
		ArchiveThreadAfter() protocol.Duration 
		BlockedWords() []string                
		Join() GroupSettings                   
	}
	GroupSettings_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	
	}

)

type (
		GroupSettings_Service_Count_Request interface { 
		GroupID() [16]byte 
	
	}
	GroupSettings_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)



type (
	GroupSettings_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	
	
	}
	GroupSettings_Service_Get_Response interface {
		GroupSettings
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)

