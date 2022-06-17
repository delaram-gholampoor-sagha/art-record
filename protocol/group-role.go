package protocol

// GroupRole indicate the domain record data fields.
type GroupRole interface {
	GroupRoleID() [16]byte                 // quiddity domain
	GroupID() [16]byte                     // group domain
	AccessControl() protocol.AccessControl //
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type GroupRole_StorageServices interface {
	Save(gr GroupRole) protocol.Error

	Count(groupRoleID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupRoleID [16]byte, versionOffset uint64) (gr GroupRole, err protocol.Error)
	Last(groupRoleID [16]byte) (gr GroupRole, numbers uint64, err protocol.Error)

	FindByGroupID(groupID [16]byte, offset, limit uint64) (groupRoleIDs [][16]byte, numbers uint64, err protocol.Error)
}
