package protocol

type StaffOvertimeStatus interface {
	StaffID() [16]byte            // staff domain
	Day() utc.DayElapsed          //
	Status() StaffOvertime_Status //
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type StaffOvertimeStatus_StorageServices interface {
	Save(gs StaffOvertimeStatus) protocol.Error

	Count(staffID [16]byte) (numbers uint64, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (gs StaffOvertimeStatus, err protocol.Error)
	Last(staffID [16]byte) (gs StaffOvertimeStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status StaffOvertime_Status, offset, limit uint64) (staffIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type StaffOvertime_Status Quiddity_Status

const (
	StaffOvertime_Status_NeedManagerApprove = StaffOvertime_Status(Quiddity_Status_FreeFlag << iota)
	StaffOvertime_Status_ManagerApprove
)
