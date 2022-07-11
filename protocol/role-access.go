package protocol

// RoleAccess indicate the domain record data fields.
type RoleAccess interface {
	RoleID() [16]byte                      // role domain
	AccessControl() protocol.AccessControl //
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type RoleAccess_StorageServices interface {
	Save(ra RoleAccess) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (ra RoleAccess, err protocol.Error)
	


}

// Get				GetRoleAccess
// GetRecursively	GetRoleAccessRecursively to get parent role access if the role has no access. bad service??
// Revoke			service to tell all node change role access control immediately


type (
	RoleAccess_Service_Register_Request interface{
		RoleID() [16]byte                      
	  AccessControl() protocol.AccessControl 
	}

	RoleAccess_Service_Register_Response interface{
	  Nv() protocol.NumberOfVersion 
	}
	
	RoleAccess_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	RoleAccess_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	RoleAccess_Service_Get_Request interface{
		RoleID() [16]byte
		VersionOffset() uint64
	}
	
	RoleAccess_Service_Get_Response interface{
		RoleAccess
		Nv() protocol.NumberOfVersion
	}
	
	RoleAccess_Service_Last_Request interface{
		RoleID() [16]byte
	}
	
	RoleAccess_Service_Last_Response interface{
		RoleAccess
		Nv() protocol.NumberOfVersion
	}
)