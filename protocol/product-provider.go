/* For license and copyright information please see LEGAL file in repository */

package art

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
	Save(pr ProductProvider) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pr ProductProvider, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductProvider_Service_Register_Request interface {
		ProductID() [16]byte 
  	RoleID() [16]byte    
		Percent() uint64
	}

	ProductProvider_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductProvider_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductProvider_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	ProductProvider_Service_Get_Request interface {
		ProductID() [16]byte
		versionOffset() uint64
	
	}
	
	ProductProvider_Service_Get_Response1 = ProductProvider
	ProductProvider_Service_Get_Response2 = protocol.NumberOfVersion
	
)