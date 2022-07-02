package protocol

type DepartmentStatus interface {
	DepartmentID() [16]byte    // department domain
	Status() Department_Status //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // user-request domain
}

type DepartmentStatus_StorageServices interface {
	Save(ds DepartmentStatus) (numbers uint64, err protocol.Error)

	Count(departmentID [16]byte) (numbers uint64, err protocol.Error)
	Get(departmentID [16]byte, versionOffset uint64) (ds DepartmentStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status Department_Status, offset, limit uint64) (departmentIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type Department_Status Quiddity_Status

const (
// Department_Status_ Department_Status = (Quiddity_Status_FreeFlag << iota)
)
