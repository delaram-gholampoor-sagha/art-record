/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductContent indicate the domain record data fields.
type ProductContent interface {
	ProductID() [16]byte // product domain
	ContentID() [16]byte // content domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductContent_StorageServices interface {
	Save(pc ProductContent) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pc ProductContent, nv protocol.NumberOfVersion, err protocol.Error)

	FindByContent(contentID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductContent_Service_Register_Request interface {
		ProductID() [16]byte 
		ContentID() [16]byte  
	}

	ProductContent_Service_Register_Response = protocol.NumberOfVersion
)

type (
		ProductContent_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductContent_Service_Count_Response = protocol.NumberOfVersion
)


type (
	ProductContent_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}

	ProductContent_Service_Get_Response1 = ProductContent
	ProductContent_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (
	ProductContent_Service_FindByContent_Request interface {
		ContentID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductContent_Service_FindByContent_Response1 interface {
		ProductIDs() [][16]byte
	}
	
	ProductContent_Service_FindByContent_Response2 = protocol.NumberOfVersion
	
)
