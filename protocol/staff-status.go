package protocol

import "../libgo/protocol"

type StaffStatus interface {
	StaffID() [16]byte    // user domain
	Status() Staff_Status //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type StaffStatus_StorageServices interface {
	Save(ss StaffStatus) protocol.Error

	Count(staffID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (ss StaffStatus, err protocol.Error)

	FilterByStatus(status Staff_Status, offset, limit uint64) (staffIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// protocol.EventTarget
}

type Staff_Status Quiddity_Status

const (
	Staff_Status_Hire = Staff_Status(Quiddity_Status_FreeFlag << iota)
	Staff_Status_Fire
	Staff_Status_End // death, end of contract, ...
	// Staff_Status_Active   // Shift
	// Staff_Status_Inactive // Shift
)

type (
	StaffStatus_Service_Register_Request interface {
		StaffID() [16]byte
		Status() Staff_Status
	}

	StaffStatus_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	StaffStatus_Service_Count_Request interface {
		StaffID() [16]byte
	}

	StaffStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	StaffStatus_Service_Get_Request interface {
		StaffID() [16]byte
		VersionOffset() uint64
	}

	StaffStatus_Service_Get_Response interface {
		StaffStatus
	}

	StaffStatus_Service_FilterByStatus_Request interface {
		StaffStatus
		Offset() uint64
		Limit() uint64
	}

	StaffStatus_Service_FilterByStatus_Response interface {
		StaffIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)