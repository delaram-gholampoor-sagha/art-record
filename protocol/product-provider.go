package protocol

import "../libgo/protocol"

// ProductProvider or product operator is to supply a particular service.
type ProductProvider interface {
	ProductID() [16]byte // product-status domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

// related UserIDs by RoleID use to get Coordinate(domain)

type ProductProvider_StorageServices interface {
	Save(pr ProductProvider) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pr ProductProvider, numbers uint64, err protocol.Error)
}

type (
	ProductProvider_Service_Register_Request interface {
		ProductID() [16]byte 
  	RoleID() [16]byte    
		Percent() uint64
	}
	ProductProvider_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductProvider_Service_Count_Request interface {
		ProductID() [16]byte
	}
	ProductProvider_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	
	}
	
	ProductProvider_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	
	}
	ProductProvider_Service_Get_Response interface {
		ProductProvider
		Nv() protocol.NumberOfVersion
	}
	
)