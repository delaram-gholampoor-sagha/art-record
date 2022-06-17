package protocol

// Role indicate the domain record data fields.
type Role interface {
	RoleID() [16]byte       // group domain
	ParentID() [16]byte     // RoleID, Exist if it isn't top of the organization.
	DepartmentID() [16]byte // department domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type Role_StorageServices interface {
	Save(r Role) (err protocol.Error)

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (r Role, err protocol.Error)
	Last(roleID [16]byte) (r Role, numbers uint64, err protocol.Error)

	GetIDs(offset, limit uint64) (roleIDs [][16]byte, numbers uint64, err protocol.Error)

	FindByDepartmentID(departmentID [16]byte, offset, limit uint64) (roleIDs [][16]byte, numbers uint64, err protocol.Error)
}
