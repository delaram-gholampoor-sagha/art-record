package protocol

import "../libgo/protocol"

// RoleLocation indicate the domain record data fields.
// Let staffs register enter and exit time by their mobiles in specific location.
type RoleLocation interface {
	RoleID() [16]byte             // role domain
	BuildingLocationID() [16]byte // building-location domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type RoleLocation_StorageServices interface {
	Save(rl RoleLocation) protocol.Error

	Count(roleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rl RoleLocation, err protocol.Error)

}

type (
	RoleLocation_Service_Register_Request interface{
		RoleID() [16]byte             
  	BuildingLocationID() [16]byte 
	}


		RoleLocation_Service_Register_Response interface{
     Nv() protocol.NumberOfVersion
	}
	
	RoleLocation_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	RoleLocation_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	RoleLocation_Service_Get_Request interface{
		RoleID() [16]byte
		VersionOffset() uint64
	}
	
	RoleLocation_Service_Get_Response interface{
		RoleLocation
		Nv() protocol.NumberOfVersion
	}
	
	RoleLocation_Service_Last_Request interface{
		RoleID() [16]byte
	}
	
	RoleLocation_Service_Last_Response interface{
		RoleLocation
		Nv() protocol.NumberOfVersion
	}
)