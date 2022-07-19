/* For license and copyright information please see LEGAL file in repository */

package art

// UserDelegate indicate the domain record data fields.
type UserDelegate interface {
	GivenUserID() protocol.UserID  // user domain
	GottenUserID() protocol.UserID // user domain
	OrgID() protocol.UserID        // user domain
	RoleID() [16]byte              // role domain
	Status() UserDelegate_Status   //
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type UserDelegate_StorageServices interface {
  Save(ud UserDelegate) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(givenUserID, gottenUserID, roleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(givenUserID, gottenUserID, roleID [16]byte, vo protocol.VersionOffset) (ud UserDelegate, nv protocol.NumberOfVersion, err protocol.Error)

	ListRoles(givenUserID, gottenUserID [16]byte, offset, limit uint64) (roleIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListOrgs(givenUserID [16]byte, offset, limit uint64) (orgIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListUsers(roleID [16]byte, offset, limit uint64) (gottenUserIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListGottenDelegate(gottenUserID [16]byte, offset, limit uint64) (givenUserIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

}

type UserDelegate_Status Quiddity_Status

const (
// UserDelegate_Status_ = UserDelegate_Status(Quiddity_Status_FreeFlag << iota)
)

// Service: delegate user can delegate its role to other. blocked this service by default.


type (

		UserDelegate_Service_Register_Request interface{
		GivenUserID() protocol.UserID  
   	GottenUserID() protocol.UserID 
  	OrgID() protocol.UserID        
  	RoleID() [16]byte             
	  Status() UserDelegate_Status   
	}
	
	UserDelegate_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
)

type (
		UserDelegate_Service_Count_Request interface{
		GivenUserID()  [16]byte
	  GottenUserID() [16]byte
		RoleID() [16]byte
		Vo() protocol.VersionOffset
	}
	

	UserDelegate_Service_Count_Response interface{
   	NumberOfVersion() protocol.NumberOfVersion
	}

)

type (

		UserDelegate_Service_Get_Request interface{
		GivenUserID()  [16]byte
	  GottenUserID() [16]byte
		RoleID() [16]byte
	}
	

	UserDelegate_Service_Get_Response interface{
		Ud() UserDelegate
		NumberOfVersion() protocol.NumberOfVersion
	}
)


type (
		UserDelegate_Service_ListRoles_Request interface{
		GivenUserID() [16]byte
		GottenUserID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	
	UserDelegate_Service_ListRoles_Response interface{
	  RoleIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)

type (

	UserDelegate_Service_ListOrgs_Request interface{
		GivenUserID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListOrgs_Response interface{
	  OrgIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}

)



type (
	UserDelegate_Service_ListGottenDelegate_Request interface{
		GottenUserID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListGottenDelegate_Response interface{
	  GivenUserIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)

type (
	UserDelegate_Service_ListUsers_Request interface{
		RoleID() [16]byte
    Offset() uint64
    Limit() uint64
	}
	
	UserDelegate_Service_ListUsers_Response interface{
	  GottenUserIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)