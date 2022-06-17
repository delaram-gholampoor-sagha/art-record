package protocol

type StaffStatus interface {
	StaffID() [16]byte    // user domain
	Status() Staff_Status //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type StaffStatus_StorageServices interface {
	Save(ss StaffStatus) protocol.Error

	Count(staffID [16]byte) (numbers uint64, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (ss StaffStatus, err protocol.Error)
	Last(staffID [16]byte) (ss StaffStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Staff_Status, offset, limit uint64) (staffIDs [][16]byte, numbers uint64, err protocol.Error)
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
