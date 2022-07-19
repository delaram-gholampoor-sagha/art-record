/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// GroupAssociatedStatus indicate the domain record data fields
type GroupAssociatedStatus interface {
	GroupID() [16]byte              // group domain
	UserID() [16]byte               // user domain
	Status() GroupAssociated_Status //
	Time() protocol.Time            // save time
	RequestID() [16]byte            // user-request domain
}

type GroupAssociatedStatus_StorageServices interface {
	Save(gn GroupAssociatedStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gn GroupAssociatedStatus, nv protocol.NumberOfVersion, err protocol.Error)
}


type GroupAssociated_Status uint8

const (
	GroupAssociated_Status_Unset GroupAssociated_Status = iota
	GroupAssociated_Status_Joined
	GroupAssociated_Status_Left
	GroupAssociated_Status_Blocked
	GroupAssociated_Status_Muted
)



// TODO::: user send last time of active state record plus its ID and optional invited user id as invite code.

type (
	GroupAssociatedStatus_Service_Register_Request interface {
		GroupID() [16]byte              
  	UserID() [16]byte               
	  Status() GroupAssociated_Status  
	}
	GroupAssociatedStatus_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
		GroupAssociatedStatus_Service_Count_Request interface { 
		GroupID() [16]byte
	}
	GroupAssociatedStatus_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (

	GroupAssociatedStatus_Service_Get_Request interface { 
		GroupID() [16]byte
		VersionOffset() uint64
	}
	GroupAssociatedStatus_Service_Get_Response interface {
		GroupAssociatedStatus
		NumberOfVersion() protocol.NumberOfVersion
	}

)