package protocol


// UserRole indicate the domain record data fields.
type UserRole interface {
	UserID() [16]byte        // user domain
	RoleID() [16]byte        // role domain
	Status() UserRole_Status //
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type UserRole_StorageServices interface {
	Save(ur UserRole) (numbers uint64, err protocol.Error)

	Count(userID, roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(userID, roleID [16]byte, versionOffset uint64) (ur UserRole, numbers uint64, err protocol.Error)

	ListRoles(userID [16]byte, offset, limit uint64) (roleIDs [][16]byte, numbers uint64, err protocol.Error)
	ListUsers(roleID [16]byte, offset, limit uint64) (userIDs [][16]byte, numbers uint64, err protocol.Error)
}

type UserRole_Status uint8

const (
	UserRole_Status_Unset UserRole_Status = iota
	UserRole_Status_Active
	UserRole_Status_Inactive
	UserRole_Status_Deleted
	UserRole_Status_Revoked
)
