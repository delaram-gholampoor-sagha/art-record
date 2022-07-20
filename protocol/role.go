/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// Role indicate the domain record data fields.
type Role interface {
	RoleID() [16]byte       // group domain
	ParentID() [16]byte     // RoleID, Exist if it isn't top of the organization.
	DepartmentID() [16]byte // department domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type Role_StorageServices interface {
	Save(r Role) (err protocol.Error)

	Count(roleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (r Role, nv protocol.NumberOfVersion ,err protocol.Error)
	

	GetIDs(offset, limit uint64) (roleIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	FindByDepartmentID(departmentID [16]byte, offset, limit uint64) (roleIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	Role_Service_Register_Request interface{      
  	ParentID() [16]byte     
	  DepartmentID() [16]byte 
	}

	Role_Service_Register_Response1 interface{
			RoleID() [16]byte
	}

	Role_Service_Register_Response2 = protocol.NumberOfVersion

)

type (
	Role_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	Role_Service_Count_Response = protocol.NumberOfVersion
)


type (
	Role_Service_Get_Request interface{
		RoleID() [16]byte
		versionOffset() uint64
	}
	
	Role_Service_Get_Response1 = Role

	Role_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	Role_Service_GetIDs_Request interface{
		Offset() uint64
		Limit() uint64
	}
	
	Role_Service_GetIDs_Response1 interface{
		RoleIDs() [][16]byte 
	}

	Role_Service_GetIDs_Response2 = protocol.NumberOfVersion
	
)


type (
	Role_Service_FindByDepartmentID_Request interface{
		DepartmentID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	Role_Service_FindByDepartmentID_Response1 interface{
		RoleIDs() [][16]byte
	}

	Role_Service_FindByDepartmentID_Response2 = protocol.NumberOfVersion
)