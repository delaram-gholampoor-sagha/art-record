package protocol

import "../libgo/protocol"

// RoleStatus indicate the domain record data fields.
type RoleStatus interface {
	RoleID() [16]byte    // role domain
	Status() Role_Status //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type RoleStatus_StorageServices interface {
	Save(rs RoleStatus) protocol.Error

	Count(roleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rs RoleStatus, err protocol.Error)


	FilterByStatus(status Role_Status, offset, limit uint64) (roleIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// protocol.EventTarget
}

type Role_Status Quiddity_Status

const (
	Role_Status_Recruit = Role_Status(Quiddity_Status_FreeFlag << iota) // Recruitment in progress
)

type (
	RoleStatus_Service_Register_Request interface{
		RoleID() [16]byte    
  	Status()	RoleStatus
	}

		RoleStatus_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	RoleStatus_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	RoleStatus_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	RoleStatus_Service_Get_Request interface{
		RoleID() [16]byte
		VersionOffset() uint64
	}
	
	RoleStatus_Service_Get_Response interface{
	  Status()	RoleStatus
	}
	

	
	RoleStatus_Service_FilterByStatus_Request interface{
		Status() RoleStatus
		Offset() uint64
		Limit() uint64
	}
	
	RoleStatus_Service_FilterByStatus_Response interface{
		RoleIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)