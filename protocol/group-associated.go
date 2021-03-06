/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// GroupAssociatedStatus indicate the domain record data fields
type GroupAssociated interface {
	GroupID() [16]byte   // group domain
	UserID() [16]byte    // user domain
	JoinBy() [16]byte    // user domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}


type GroupAssociated_StorageServices interface {
	Save(gn GroupAssociated) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gn GroupAssociated, nv protocol.NumberOfVersion, err protocol.Error)
}


// TODO::: user send last time of active state record plus its ID and optional invited user id as invite code.


type (
	GroupAssociated_Service_Register_Request interface {
		GroupID() [16]byte  
		UserID() [16]byte    
		JoinBy() [16]byte    
	}

	GroupAssociated_Service_Register_Response = protocol.NumberOfVersion

)

type (
	GroupAssociated_Service_Count_Request interface { 
		GroupID() [16]byte
	}

	GroupAssociated_Service_Count_Response = protocol.NumberOfVersion

)

type (
	GroupAssociated_Service_Get_Request interface { 
		GroupID() [16]byte
		versionOffset() uint64
	}
	GroupAssociated_Service_Get_Response1 =	GroupAssociated

	GroupAssociated_Service_Get_Response2 = protocol.NumberOfVersion

)