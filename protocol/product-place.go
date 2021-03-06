/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductPlace indicate the domain record data fields.
// ProductPlace indicate service provides in specific place e.g. stations(bus, train, airport, ...), ...
type ProductPlace interface {
	ProductID() [16]byte          // product domain.
	BuildingLocationID() [16]byte // building-location domain. use also for Coordinate domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type ProductPlace_StorageServices interface {
	Save(pp ProductPlace) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pp ProductPlace, nv protocol.NumberOfVersion, err protocol.Error)

	FindByBuildingLocation(buildingLocationID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductPlace_Service_Register_Request interface {
		ProductID() [16]byte
		BuildingLocationID() [16]byte
	}

	ProductPlace_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductPlace_Service_Count_Request interface {
		ProductID() [16]byte
	}
	ProductPlace_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	ProductPlace_Service_Get_Request interface {
		ProductID() [16]byte
		versionOffset() uint64
	}

	ProductPlace_Service_Get_Response1 = ProductPlace
	ProductPlace_Service_Get_Response2 = protocol.NumberOfVersion	
)



type (
	ProductPlace_Service_FindByBuildingLocation_Request interface {
		BuildingLocationID() [16]byte
		Offset() uint64
		limit() uint64
	}
	ProductPlace_Service_FindByBuildingLocation_Response1 interface {
		ProductIDs() [][16]byte
	}

	ProductPlace_Service_FindByBuildingLocation_Response2 = protocol.NumberOfVersion
	
)
