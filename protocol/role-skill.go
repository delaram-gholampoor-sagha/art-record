package protocol


// RoleSkill indicate the domain record data fields.
type RoleSkill interface {
	RoleID() [16]byte // role domain

	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
