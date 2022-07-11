package protocol

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
	Get(userID [16]byte, versionOffset uint64) (sh StaffHire, err protocol.Error)
	Last(userID [16]byte) (sh StaffHire, nv protocol.NumberOfVersion, err protocol.Error)

	FindByRoleID(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	// protocol.EventTarget
}

type (
	StaffHire_Service_Register_Request interface {
		UserID() [16]byte    
  	RoleID() [16]byte    
	}

		StaffHire_Service_Register_Response interface {
	  	Nv() protocol.NumberOfVersion
	}

	StaffHire_Service_Count_Request interface {
		StaffID() [16]byte
	}

	StaffHire_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	StaffHire_Service_Get_Request interface {
		StaffID() [16]byte
		VersionOffset() uint64
	}

	StaffHire_Service_Get_Response interface {
		StaffHire
	}

	StaffHire_Service_Last_Request interface {
		StaffID() [16]byte
	}

	StaffHire_Service_Last_Response interface {
		StaffHire
		Nv() protocol.NumberOfVersion
	}

	StaffHire_Service_FindByRoleID_Request interface {
		RoleID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	StaffHire_Service_FindByRoleID_Response interface {
		StaffIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)