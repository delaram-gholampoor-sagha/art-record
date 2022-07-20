/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductArea indicate the domain record data fields.
// ProductArea is proper or appropriate location indicate specific area not specific location e.g. Taxi, ...
type ProductArea interface {
	ProductID() [16]byte // product domain
	AreaID() [16]byte    // area domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductArea_StorageServices interface {
	Save(pa ProductArea) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pa ProductArea, nv protocol.NumberOfVersion, err protocol.Error)

	FindByArea(AreaID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductArea_Service_Register_Request interface {
		ProductID() [16]byte 
		AreaID() [16]byte    
	}

	ProductArea_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductArea_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductArea_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	ProductArea_Service_Get_Request interface {
		ProductID() [16]byte    
		versionOffset() uint64
	}

	ProductArea_Service_Get_Response = ProductArea
	ProductArea_Service_Get_Response2 = protocol.NumberOfVersion

)



type (
	ProductArea_Service_FindByArea_Request interface {
		AreaID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductArea_Service_FindByArea_Response1 interface {
		ProductIDs() [][16]byte
	}

	ProductArea_Service_FindByArea_Response2 = protocol.NumberOfVersion
)