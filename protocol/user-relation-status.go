/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// UserRelationStatus indicate the domain record data fields.
type UserRelationStatus interface {
	UserID() [16]byte            // user domain
	SideID() [16]byte            // user domain
	Status() UserRelation_Status //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type UserRelationStatus_StorageServices interface {
	Save(gs UserRelationStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(userID, sideID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID, sideID [16]byte, versionOffset uint64) (gs UserRelationStatus, nv protocol.NumberOfVersion, err protocol.Error)

	ListSides(userID [16]byte, offset, limit uint64) (sideIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListUsers(sideID [16]byte, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	FilterByStatus(userID [16]byte, status UserRelation_Status, offset, limit uint64) (sideIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	protocol.EventTarget
}

type UserRelation_Status Quiddity_Status

const (
	UserRelation_Status_Blocked = UserRelation_Status(Quiddity_Status_FreeFlag << iota)
	UserRelation_Status_Muted
)


type (
	UserRelationStatus_Service_Register_Request interface{
		UserID() [16]byte            
	  SideID() [16]byte            
  	Status() UserRelation_Status 
	}

	UserRelationStatus_Service_Register_Response = protocol.NumberOfVersion

)


type (
	UserRelationStatus_Service_Count_Request interface{
		UserID() [16]byte
	}

	UserRelationStatus_Service_Count_Response = protocol.NumberOfVersion

)


type (
	UserRelationStatus_Service_Get_Request interface{
		UserID() [16]byte
		versionOffset() uint64
	}

	UserRelationStatus_Service_Get_Response1 = UserRelationStatus
	UserRelationStatus_Service_Get_Response2 = protocol.NumberOfVersion
)

type (
	UserRelationStatus_Service_ListSides_Request interface{
		UserID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	UserRelationStatus_Service_ListSides_Response1 interface{
		SideIDs() [][16]byte
	}

	UserRelationStatus_Service_ListSides_Response2 = protocol.NumberOfVersion
)

type (
	UserRelationStatus_Service_ListUsers_Request interface{
		SideID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	UserRelationStatus_Service_ListUsers_Response1 interface{
		UserIDs() [][16]byte
	}

	UserRelationStatus_Service_ListUsers_Response2 = protocol.NumberOfVersion
)

type (

	UserRelationStatus_Service_FilterByStatus_Request interface{
		UserID() [16]byte
		Status() UserRelationStatus
		Offset() uint64
		Limit() uint64
	}

	UserRelationStatus_Service_FilterByStatus_Response1 interface{
		SideIDs() [][16]byte
	}

	UserRelationStatus_Service_FilterByStatus_Response2 = protocol.NumberOfVersion
)