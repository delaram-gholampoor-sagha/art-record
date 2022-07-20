/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// BuildingLocation can use to indicate any part of a building like warehouses, zones, halls, ...
type BuildingLocation interface {
	BuildingLocationID() [16]byte // quiddity domain. use in Coordinate domain too, to get physical location of the Location.
	BuildingID() [16]byte         // building domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}


type BuildingLocation_StorageServices interface {
	Save(c BuildingLocation) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(BuildingLocationID [16]byte , BuildingID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(BuildingLocationID [16]byte, BuildingID [16]byte , vo protocol.versionOffset) (gs BuildingLocation, nv protocol.NumberOfVersion, err protocol.Error)

	GetIDs(offset, limit uint64) (BuildingLocationIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	BuildingLocation_Service_Register_Request interface {
		BuildingLocationID() [16]byte  
		BuildingID() uint64 	
	}

	BuildingLocation_Service_Register_Response = protocol.NumberOfVersion

)

type (
	BuildingLocation_Service_Count_Request interface {
		BuildingLocationID() [16]byte  
		BuildingID() uint64 	
	}
	BuildingLocation_Service_Count_Response = protocol.NumberOfVersion

)

type (
	BuildingLocation_Service_Get_Request interface {
		BuildingLocationID() [16]byte  
		BuildingID() uint64
		versionOffset() protocol.versionOffset
	}

	BuildingLocation_Service_Get_Response1 = 	BuildingLocation
	BuildingLocation_Service_Get_Response2 = protocol.NumberOfVersion
	
)



type (
	BuildingLocation_Service_GetIDs_Request interface {
		Offset() uint64
		Limit() uint64
	}
	BuildingLocation_Service_GetIDs_Response1 interface {
		BuildingLocationIDs() [][16]byte			
	}
	
	BuildingLocation_Service_GetIDs_Response2 = protocol.NumberOfVersion 

)

