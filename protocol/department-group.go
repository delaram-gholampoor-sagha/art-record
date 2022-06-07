package protocol

type DepartmentGroup interface {
	DepartmentID() [16]byte // quiddity domain. use to get and show locale title.
	GroupID() [16]byte      // group id
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type DepartmentGroup_StorageServices interface {
	Save(dg DepartmentGroup) (err protocol.Error)
	Delete(departmentID [16]byte, groupID [16]byte) (err protocol.Error)

	ListGroupIDs(departmentID [16]byte, offset, limit uint64) (groupIDs [][16]byte, numbers uint64, err protocol.Error)
}
