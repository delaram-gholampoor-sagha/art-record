/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// GroupRelated indicate the domain record data fields.
type GroupRelated interface {
	GroupID() [16]byte        // group domain
	RelatedGroupID() [16]byte // group domain. Use to categorize and graph groups
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type GroupRelated_StorageServices interface {
	Save(gr GroupRelated) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(GroupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(GroupID [16]byte, versionOffset uint64) (gr GroupRelated , nv protocol.NumberOfVersion, err protocol.Error)

	FindByRelatedGroupID(groupID [16]byte, offset, limit uint64) (RelatedGroupID [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	GroupRelated_Service_Register_Request interface {
	 GroupID() [16]byte        
	 RelatedGroupID() [16]byte              
	}

	GroupRelated_Service_Register_Response = protocol.NumberOfVersion

)

type (
	GroupRelated_Service_Count_Request interface { 
		GroupID() [16]byte
	}

	GroupRelated_Service_Count_Response = protocol.NumberOfVersion

)


type (
	GroupRelated_Service_Get_Request interface { 
		GroupID() [16]byte
		versionOffset() uint64
	}

	GroupRelated_Service_Get_Response1 =	GroupRelated
	GroupRelated_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (
	GroupRelated_Service_FindByGroupID_Request interface { 
		GroupID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	GroupRelated_Service_FindByGroupID_Response interface {
		GroupRelatedIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)
