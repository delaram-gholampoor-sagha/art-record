package protocol

type BuildingLocationStatus interface {
	BuildingLocationID() [16]byte    // building-location domain
	Status() BuildingLocation_Status //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

type BuildingLocationStatus_StorageServices interface {
	Save(bs BuildingLocationStatus) protocol.Error

	Count(buildingLocationID [16]byte) (numbers uint64, err protocol.Error)
	Get(buildingLocationID [16]byte, versionOffset uint64) (bs BuildingLocationStatus, err protocol.Error)
	Last(buildingLocationID [16]byte) (bs BuildingLocationStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status BuildingLocation_Status, offset, limit uint64) (buildingLocationIDs [][16]byte, err protocol.Error)
	// protocol.EventTarget
}

type BuildingLocation_Status Quiddity_Status

const (
// BuildingLocation_Status_ = BuildingLocation_Status(Quiddity_Status_FreeFlag << iota)
)
