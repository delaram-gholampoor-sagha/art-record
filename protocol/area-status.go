package protocol



type AreaStatus interface {
	AreaID() [16]byte    // quiddity domain
	Status() Area_Status //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type AreaStatus_StorageServices interface {
	Save(as AreaStatus) protocol.Error

	Count(areaID [16]byte) (numbers uint64, err protocol.Error)
	Get(areaID [16]byte, versionOffset uint64) (as AreaStatus, err protocol.Error)
	Last(areaID [16]byte) (as AreaStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Area_Status, offset, limit uint64) (areaIDs [][16]byte, err protocol.Error)
	// protocol.EventTarget
}

type Area_Status Quiddity_Status

const (
	Area_Status_NeighborsConflict = Area_Status(Quiddity_Status_FreeFlag << iota)
)
