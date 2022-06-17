package protocol

// GroupAssociatedStatus indicate the domain record data fields
type GroupAssociated interface {
	GroupID() [16]byte   // group domain
	UserID() [16]byte    // user domain
	JoinBy() [16]byte    // user domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
