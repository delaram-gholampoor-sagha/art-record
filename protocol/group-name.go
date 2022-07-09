package protocol

// GroupName indicate the domain record data fields.
// It is not the group title as display name, It is the unique group name.
type GroupName interface {
	GroupID() [16]byte        // group domain
	Name() string             // It is not replace of group ID. It usually use to find the group in more human friendly manner.
	Status() GroupName_Status //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type GroupName_StorageServices interface {
	Save(gn GroupName) (numbers uint64, err protocol.Error)

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gn GroupName, numbers uint64, err protocol.Error)

	FindByName(name string) (groupID [16]byte, err protocol.Error)
}



// GroupName_Status indicate GroupName record status
type GroupName_Status uint8

// GroupName status
const (
	GroupName_Status_Unset    GroupName_Status = 0
	GroupName_Status_Register GroupName_Status = (1 << iota)
	GroupName_Status_Remove
	GroupName_Status_Blocked
	GroupName_Status_Reserved
)

type (
	
	GroupName_Service_Register_Request interface {
		GroupID() [16]byte                 
		Name() string
		Status() GroupName_Status  
	}
	GroupName_Service_Register_Response interface {
		Numbers() uint64
	}

	GroupName_Service_Count_Request interface { 
		GroupID() [16]byte
	}
	GroupName_Service_Count_Response interface {
		Numbers() uint64
	}
	GroupName_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	}
	GroupName_Service_Get_Response interface {
		GroupName
		Numbers() uint64
	}
	GroupName_Service_FindByName_Request interface { 
		Name() string
	}
	GroupName_Service_FindByName_Response interface {
		GroupID() [16]byte
	}
)
