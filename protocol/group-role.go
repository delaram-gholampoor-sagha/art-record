package protocol

// GroupRole indicate the domain record data fields.
type GroupRole interface {
	GroupRoleID() [16]byte                 // quiddity domain
	GroupID() [16]byte                     // group domain
	AccessControl() protocol.AccessControl //
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type GroupRole_StorageServices interface {
	Save(gr GroupRole) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(groupRoleID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(groupRoleID [16]byte, versionOffset uint64) (gr GroupRole, nv protocol.NumberOfVersion, err protocol.Error)

	FindByGroupID(groupID [16]byte, offset, limit uint64) (groupRoleIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	GroupRole_Service_Register_Request interface {
		GroupRoleID() [16]byte               
		GroupID() [16]byte                     
		AccessControl() protocol.AccessControl 	             
	}
	GroupRole_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	GroupRole_Service_Count_Request interface { 
		GroupRoleID() [16]byte
	}
	GroupRole_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	GroupRole_Service_Get_Request interface { 
		GroupRoleID() [16]byte
		VersionOffset() uint64
	}
	GroupRole_Service_Get_Response interface {
		GroupRole
		Nv() protocol.NumberOfVersion
	}
	GroupRole_Service_FindByGroupID_Request interface { 
		GroupID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	GroupRole_Service_FindByGroupID_Response interface {
		GroupRoleIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)
