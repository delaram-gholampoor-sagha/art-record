package protocol

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
	Get(BuildingLocationID [16]byte, BuildingID [16]byte , vo protocol.VersionOffset) (gs BuildingLocation, nv protocol.NumberOfVersion, err protocol.Error)

	GetIDs(offset, limit uint64) (BuildingLocationIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
		BuildingLocation_Service_Register_Request interface {
		BuildingLocationID() [16]byte  
		BuildingID() uint64 	
	}
	BuildingLocation_Service_Register_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
		BuildingLocation_Service_Count_Request interface {
		BuildingLocationID() [16]byte  
		BuildingID() uint64 

			
	}
	BuildingLocation_Service_Count_Response interface {
		NumberOfVersion() protocol.NumberOfVersion
			
	}

)

type (
	BuildingLocation_Service_Get_Request interface {
		BuildingLocationID() [16]byte  
		BuildingID() uint64
		Vo() protocol.VersionOffset
	}
	BuildingLocation_Service_Get_Response interface {
		BuildingLocation
		NumberOfVersion() protocol.NumberOfVersion
			
	}
	
)



type (
	BuildingLocation_Service_GetIDs_Request interface {
		Offset() uint64
		Limit() uint64
	}
	BuildingLocation_Service_GetIDs_Response interface {
		BuildingLocationIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
			
	}

)

