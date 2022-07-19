/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// GroupAssociatedRole indicate the domain record data fields.
type GroupAssociatedRole interface {
	GroupID() [16]byte     // group domain
	UserID() [16]byte      // user domain
	GroupRoleID() [16]byte // group-role domain
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}


type GroupAssociatedRole_StorageServices interface {
	Save(gn GroupAssociatedRole) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gn GroupAssociatedRole, nv protocol.NumberOfVersion, err protocol.Error)
}



// TODO::: user send last time of active state record plus its ID and optional invited user id as invite code.


type (
	GroupAssociatedRole_Service_Register_Request interface {
		GroupID() [16]byte     
  	UserID() [16]byte      
  	GroupRoleID() [16]byte  
	}
	GroupAssociatedRole_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
		GroupAssociatedRole_Service_Count_Request interface { 
		GroupID() [16]byte
	}
	GroupAssociatedRole_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
)

type (
	GroupAssociatedRole_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	}
	GroupAssociatedRole_Service_Get_Response interface {
		GroupAssociatedRole
		NumberOfVersion() protocol.NumberOfVersion
	}

)