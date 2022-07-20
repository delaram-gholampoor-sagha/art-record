/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type StaffHire interface {
	UserID() [16]byte    // user domain
	RoleID() [16]byte    // role status
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type StaffHire_StorageServices interface {
	Save(sh StaffHire) protocol.Error

	Count(userID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (sh StaffHire , nv protocol.NumberOfVersion, err protocol.Error)
	

	FindByRoleID(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	// protocol.EventTarget
}


type (
	StaffHire_Service_Register_Request interface {
		UserID() [16]byte    
  	RoleID() [16]byte    
	}

	StaffHire_Service_Register_Response = protocol.NumberOfVersion

)

type (
	StaffHire_Service_Count_Request interface {
		StaffID() [16]byte
	}

	StaffHire_Service_Count_Response = protocol.NumberOfVersion
)


type (
		StaffHire_Service_Get_Request interface {
		StaffID() [16]byte
		versionOffset() uint64
	}

	StaffHire_Service_Get_Response1 = StaffHire
	StaffHire_Service_Get_Response2 = protocol.NumberOfVersion
	
)

type (
	StaffHire_Service_FindByRoleID_Request interface {
		RoleID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	StaffHire_Service_FindByRoleID_Response1 interface {
		StaffIDs() [][16]byte
	}

	StaffHire_Service_FindByRoleID_Response2 = protocol.NumberOfVersion 
)