package protocol


// UserDelegate indicate the domain record data fields.
type UserDelegate interface {
	UserID() protocol.UserID         // user domain
	OrgID() protocol.UserID          //
	DelegateUserID() protocol.UserID // user domain
	RoleID() [16]byte                // role domain
	Status() UserDelegate_Status     //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

type UserDelegate_StorageServices interface {
	Save(ud UserDelegate) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(userID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID [16]byte, vo protocol.VersionOffset) (ud UserDelegate, nv protocol.NumberOfVersion, err protocol.Error)

	CountStatus(userID, orgID ,delegateUserID,  roleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	GetStatus(userID, orgID ,delegateUserID,  roleID [16]byte ,vo protocol.VersionOffset) (ud UserDelegate, nv protocol.NumberOfVersion, err protocol.Error)

	ListOrgs(userID [16]byte, offset, limit uint64) (orgIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListRoles(userID [16]byte, offset, limit uint64) (roleIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListGivenDelegate(userID [16]byte, offset, limit uint64) (delegateUserIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListGottenDelegate(delegateUserID [16]byte, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListUsers(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type UserDelegate_Status Quiddity_Status

const (
// UserDelegate_Status_ = UserDelegate_Status(Quiddity_Status_FreeFlag << iota)
)

// Service: delegate user can delegate its role to other. blocked this service by default.

type (
	UserDelegate_Service_Register_Request interface{
		UserID() protocol.UserID         
		OrgID() protocol.UserID          
		DelegateUserID() protocol.UserID 
		RoleID() [16]byte                
		Status() UserDelegate_Status     
	}
	
	UserDelegate_Service_Register_Response interface{
		Nv() protocol.NumberOfVersion
	}
	
	UserDelegate_Service_Count_Request interface{
		UserID() [16]byte
	}
	

	UserDelegate_Service_Count_Response interface{
   	Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_Get_Request interface{
		UserID() [16]byte
		Vo() protocol.VersionOffset
	}
	

	UserDelegate_Service_Get_Response interface{
		Ud() UserDelegate
		Nv() protocol.NumberOfVersion
	}

	

	UserDelegate_Service_CountStatus_Request interface{
			UserID() [16]byte
			OrgID() [16]byte
			DelegateUserID() [16]byte
			RoleID() [16]byte
			Nv() protocol.NumberOfVersion
	}
	
	UserDelegate_Service_CountStatus_Response interface{
		Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_GetStatus_Request interface{
		UserID() [16]byte
			OrgID() [16]byte
			DelegateUserID() [16]byte
			RoleID() [16]byte
	}
	
	UserDelegate_Service_GetStatus_Response interface{
		  Ud() UserDelegate
			Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_ListOrgs_Request interface{
		UserID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListOrgs_Response interface{
	  OrgIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_ListRoles_Request interface{
			UserID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListRoles_Response interface{
  RoleIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_ListGivenDelegate_Request interface{
		UserID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListGivenDelegate_Response interface{
	 DelegateUserIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_ListGottenDelegate_Request interface{
		DelegateUserID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListGottenDelegate_Response interface{
	  UserIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}

	UserDelegate_Service_ListUsers_Request interface{
		RoleID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListUsers_Response interface{
	  UserIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
	
)