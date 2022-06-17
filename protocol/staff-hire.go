package protocol

type StaffHire interface {
	UserID() [16]byte    // user domain
	RoleID() [16]byte    // role status
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type StaffHire_StorageServices interface {
	Save(sh StaffHire) protocol.Error

	Count(userID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID [16]byte, versionOffset uint64) (sh StaffHire, err protocol.Error)
	Last(userID [16]byte) (sh StaffHire, numbers uint64, err protocol.Error)

	FindByRoleID(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)

	// protocol.EventTarget
}
