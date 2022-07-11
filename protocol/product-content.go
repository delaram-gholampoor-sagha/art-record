package protocol

import "../libgo/protocol"

// ProductContent indicate the domain record data fields.
type ProductContent interface {
	ProductID() [16]byte // product domain
	ContentID() [16]byte // content domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductContent_StorageServices interface {
	Save(pc ProductContent) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pc ProductContent, numbers uint64, err protocol.Error)

	FindByContent(contentID [16]byte, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	ProductContent_Service_Register_Request interface {
		ProductID() [16]byte 
		ContentID() [16]byte  
	}
	ProductContent_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductContent_Service_Count_Request interface {
		ProductID() [16]byte
	}
	ProductContent_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductContent_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductContent_Service_Get_Response interface {
		ProductContent
		Nv() protocol.NumberOfVersion
	}
	
	ProductContent_Service_FindByContent_Request interface {
		ContentID() [16]byte
		Offset() uint64
		Limit() uint64}
	ProductContent_Service_FindByContent_Response interface {
		ProductIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
	
)
