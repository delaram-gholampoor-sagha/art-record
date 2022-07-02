package protocol
// GroupAssociatedRole indicate the domain record data fields.
type GroupAssociatedRole interface {
	GroupID() [16]byte     // group domain
	UserID() [16]byte      // user domain
	GroupRoleID() [16]byte // group-role domain
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}
