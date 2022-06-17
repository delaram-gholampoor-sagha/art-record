package protocol

type StaffMissionStatus interface {
	StaffID() [16]byte           // staff domain
	Day() utc.DayElapsed         //
	Status() StaffMission_Status //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type StaffMissionStatus_StorageServices interface {
	Save(gs StaffMissionStatus) protocol.Error

	Count(staffID [16]byte) (numbers uint64, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (gs StaffMissionStatus, err protocol.Error)
	Last(staffID [16]byte) (gs StaffMissionStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status StaffMission_Status, offset, limit uint64) (staffIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type StaffMission_Status Quiddity_Status

const (
	StaffMission_Status_NeedManagerApprove = StaffMission_Status(Quiddity_Status_FreeFlag << iota)
	StaffMission_Status_ManagerApproved
)
