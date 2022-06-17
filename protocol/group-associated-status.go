package protocol

// GroupAssociatedStatus indicate the domain record data fields
type GroupAssociatedStatus interface {
	GroupID() [16]byte              // group domain
	UserID() [16]byte               // user domain
	Status() GroupAssociated_Status //
	Time() protocol.Time            // save time
	RequestID() [16]byte            // user-request domain
}

type GroupAssociated_Status uint8

const (
	GroupAssociated_Status_Unset GroupAssociated_Status = iota
	GroupAssociated_Status_Joined
	GroupAssociated_Status_Left
	GroupAssociated_Status_Blocked
	GroupAssociated_Status_Muted
)
