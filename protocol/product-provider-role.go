package protocol

import "../libgo/protocol"

// ProductProviderRole indicate the domain record data fields.
// ProductProviderRole or product operator is to supply a particular service.
type ProductProviderRole interface {
	ProductID() [16]byte // product domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductProviderRole_StorageServices interface {
	Save(pp ProductProviderRole) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pp ProductProviderRole, numbers uint64, err protocol.Error)

	FindByRole(roleID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

// related UserIDs by RoleID use to get Coordinate(domain)

type (
	ProductProviderRole_Service_Register_Request interface {
		ProductID() [16]byte
		RoleID() [16]byte
	}
	ProductProviderRole_Service_Register_Response interface {
		Numbers() uint64
	}
	
	ProductProviderRole_Service_Count_Request interface {
		ProductID() [16]byte
	}
	ProductProviderRole_Service_Count_Response interface {
		Numbers() uint64
	}
	
	ProductProviderRole_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductProviderRole_Service_Get_Response interface {
		ProductProviderRole
		Numbers() uint64
		
	}
	
	ProductProviderRole_Service_FindByRole_Request interface {
		RoleID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductProviderRole_Service_FindByRole_Response interface {
		ProductionIds() [][16]byte
		Numbers() uint64
	}
	
)