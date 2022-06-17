package protocol


// RoleAccess indicate the domain record data fields.
type RoleAccess interface {
	RoleID() [16]byte                      // role domain
	AccessControl() protocol.AccessControl //
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type RoleAccess_StorageServices interface {
	Save(ra RoleAccess) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (ra RoleAccess, err protocol.Error)
	Last(roleID [16]byte) (ra RoleAccess, numbers uint64, err protocol.Error)

	protocol.EventTarget
}

// Get				GetRoleAccess
// GetRecursively	GetRoleAccessRecursively to get parent role access if the role has no access. bad service??
// Revoke			service to tell all node change role access control immediately
