package protocol

// PosProvider or pos operator is to limit pos to specific staffs.
type PosProvider interface {
	PosID() [16]byte     // pos domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
