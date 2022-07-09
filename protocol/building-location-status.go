package protocol

import "../libgo/protocol"

type BuildingLocationStatus interface {
	BuildingLocationID() [16]byte    // building-location domain
	Status() BuildingLocation_Status //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}
type BuildingLocationStatus_StorageServices interface {
	Save(bs BuildingLocationStatus) (numbers uint64, err protocol.Error)

	Count(buildingLocationID [16]byte) (numbers uint64, err protocol.Error)
	Get(buildingLocationID [16]byte, versionOffset uint64) (bs BuildingLocationStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status BuildingLocation_Status, offset, limit uint64) (buildingLocationIDs [][16]byte, err protocol.Error)
	// protocol.EventTarget
}

type BuildingLocation_Status Quiddity_Status

const (
// BuildingLocation_Status_ = BuildingLocation_Status(Quiddity_Status_FreeFlag << iota)
)

type (
	BuildingLocationStatus_Service_Register_Request interface {
		BuildingLocationID() [16]byte  
		Status() BuildingLocation_Status	
	}
	
	BuildingLocationStatus_Service_Register_Response interface {
		Numbers()uint64
	}
			
	BuildingLocationStatus_Service_Count_Request interface {
		BuildingLocationID() [16]byte  
		
			
	}
	BuildingLocationStatus_Service_Count_Response interface {
		Numbers() uint64
			
	}
	BuildingLocationStatus_Service_Get_Request interface {
		BuildingLocationID() [16]byte  
		VersionOffset() uint64
	}
	BuildingLocationStatus_Service_Get_Response interface {
		BuildingLocationStatus
		Numbers() uint64		
	}

	BuildingLocationStatus_Service_FilterByStatus_Request interface {
		Status() BuildingLocation_Status
		Offset() uint64
		Limit() uint64
	}
	BuildingLocationStatus_Service_FilterByStatus_Response interface {
		BuildingLocationIDs() [][16]byte		
	}

)

