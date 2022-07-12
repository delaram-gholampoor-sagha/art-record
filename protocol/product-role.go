package protocol

import "../libgo/protocol"

// ProductRole indicate the domain record data fields.
// ProductRole limit sell to specific roles.
type ProductRole interface {
	ProductID() [16]byte // product domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductRole_StorageServices interface {
	Save(pr ProductRole) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pr ProductRole, nv protocol.NumberOfVersion, err protocol.Error)

	FindByRole(roleID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductRole_Service_Register_Request interface{
		ProductID() [16]byte
		RoleID() [16]byte
	
	}
	ProductRole_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	ProductRole_Service_Count_Request interface{
		ProductID() [16]byte
	}
	ProductRole_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	
	}
	
	ProductRole_Service_Get_Request interface{
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductRole_Service_Get_Response interface{
		ProductRole
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	ProductRole_Service_FindByRole_Request interface{
		RoleID() [16]byte
		Offset() uint64 
		Limit() uint64
	}
	ProductRole_Service_FindByRole_Response interface{
		ProductIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)




