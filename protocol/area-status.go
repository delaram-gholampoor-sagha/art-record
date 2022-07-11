package protocol

import "../libgo/protocol"

type AreaStatus interface {
	AreaID() [16]byte    // quiddity domain
	Status() Area_Status //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type AreaStatus_StorageServices interface {
	Save(as AreaStatus) protocol.Error

	Count(areaID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(areaID [16]byte, versionOffset uint64) (as AreaStatus, err protocol.Error)

	FilterByStatus(status Area_Status, offset, limit uint64) (areaIDs [][16]byte, err protocol.Error)
	// protocol.EventTarget
}

type Area_Status Quiddity_Status

const (
	Area_Status_NeighborsConflict = Area_Status(Quiddity_Status_FreeFlag << iota)
)

type (
	AreaStatus_Service_Register_Request interface {
		AreaID() [16]byte
		Status() Area_Status
	}
	AreaStatus_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	AreaStatus_Service_Count_Request interface {
		AreaID() [16]byte
	}
	AreaStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	AreaStatus_Service_Get_Request interface {
		AreaID() [16]byte
		VersionOffset() uint64
	}
	AreaStatus_Service_Get_Response interface {
		Status() AreaStatus
		Nv() protocol.NumberOfVersion
	}

	AreaStatus_Service_FilterByStatus_Request interface {
		Status() AreaStatus
		Offset() uint64
		Limit() uint64
	}
	AreaStatus_Service_FilterByStatus_Response interface {
		areaIDs() [][16]byte
	}
)
