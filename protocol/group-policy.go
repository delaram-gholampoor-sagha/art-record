/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// GroupPolicy indicate the domain record data fields.
type GroupPolicy interface {
	GroupID() [16]byte                 // group domain
	GetPolicy() GroupPolicy_Policy     // Visibility level
	SetPolicy() GroupPolicy_Policy     //
	ForwardPolicy() GroupPolicy_Policy //
	Time() protocol.Time               // save time
	RequestID() [16]byte               // user-request domain
}

type GroupPolicy_StorageServices interface {
	Save(cg GroupPolicy) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (cg GroupPolicy, nv protocol.NumberOfVersion, err protocol.Error)

	FindByUserID(ownerUserID [16]byte, offset, limit uint64) (ids [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}



type GroupPolicy_Policy uint16

const (
	GroupPolicy_Policy_Unset     GroupPolicy_Policy = 0
	GroupPolicy_Policy_JustOwner GroupPolicy_Policy = (1 << iota)
	GroupPolicy_Policy_GroupManagers
	GroupPolicy_Policy_GroupMembers
	GroupPolicy_Policy_Mentioned
	GroupPolicy_Policy_AnyOne
	// Below policies are person type authorization
	GroupPolicy_Policy_1stFriendshipCycle
	GroupPolicy_Policy_2stFriendshipCycle
	GroupPolicy_Policy_3stFriendshipCycle
	GroupPolicy_Policy_4stFriendshipCycle
	GroupPolicy_Policy_5stFriendshipCycle
	GroupPolicy_Policy_6stFriendshipCycle
)


type (
	GroupPolicy_Service_Register_Request interface {
		GroupID() [16]byte                 
		GetPolicy() GroupPolicy_Policy     
		SetPolicy() GroupPolicy_Policy    
		ForwardPolicy() GroupPolicy_Policy            
	}

	GroupPolicy_Service_Register_Response = protocol.NumberOfVersion

)

type (
	GroupPolicy_Service_Count_Request interface { 
		GroupID() [16]byte
	}

	GroupPolicy_Service_Count_Response = protocol.NumberOfVersion
	
)



type (
	GroupPolicy_Service_Get_Request interface { 
		GroupID() [16]byte
		versionOffset() uint64
	}

	GroupPolicy_Service_Get_Response1 = GroupPolicy
	GroupPolicy_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (
	GroupPolicy_Service_FindByUserID_Request interface { 
		OwnerUserID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	GroupPolicy_Service_FindByUserID_Response1 interface {
		IDs() [][16]byte
	}

	GroupPolicy_Service_FindByUserID_Response2 =  protocol.NumberOfVersion
)