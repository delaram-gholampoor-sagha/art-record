/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductGroup indicate the domain record data fields.
type ProductGroup interface {
	ProductID() [16]byte // product domain
	GroupID() [16]byte   // group domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductGroup_StorageServices interface {
	Save(pg ProductGroup) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pg ProductGroup, nv protocol.NumberOfVersion, err protocol.Error)

	FindByGroup(groupID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductGroup_Service_Register_Request interface {
		ProductID() [16]byte 
		GroupID() [16]byte  
	}
	ProductGroup_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductGroup_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductGroup_Service_Count_Response = protocol.NumberOfVersion
)



type (
	ProductGroup_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}

	ProductGroup_Service_Get_Response1 = 	ProductGroup
	ProductGroup_Service_Get_Response2 =  protocol.NumberOfVersion
	
)

type (
	
	ProductGroup_Service_FindByGroup_Request interface {
		GroupID() [16]byte
		Offset() uint64
		Limit() uint64}
	ProductGroup_Service_FindByGroup_Response1 interface {
		ProductIDs() [][16]byte
	}
	
	ProductGroup_Service_FindByGroup_Response2 = protocol.NumberOfVersion
)
