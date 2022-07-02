package protocol

type DepartmentOffice interface {
	DepartmentID() [16]byte       // department domain
	BuildingLocationID() [16]byte // building-location domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type DepartmentOffice_StorageServices interface {
	Save(do DepartmentOffice) (numbers uint64, err protocol.Error)

	Count(departmentID [16]byte) (numbers uint64, err protocol.Error)
	Get(departmentID [16]byte, versionOffset uint64) (do DepartmentOffice, numbers uint64, err protocol.Error)
}
