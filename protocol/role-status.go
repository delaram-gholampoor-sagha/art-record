package protocol

// RoleStatus indicate the domain record data fields.
type RoleStatus interface {
	RoleID() [16]byte    // role domain
	Status() Role_Status //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type RoleStatus_StorageServices interface {
	Save(rs RoleStatus) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rs RoleStatus, err protocol.Error)
	Last(roleID [16]byte) (rs RoleStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Role_Status, offset, limit uint64) (roleIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type Role_Status Quiddity_Status

const (
	Role_Status_Recruit = Role_Status(Quiddity_Status_FreeFlag << iota) // Recruitment in progress
)
