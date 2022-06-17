package protocol

// GroupInviteStatus indicate the domain record data fields.
type GroupInviteStatus interface {
	GroupID() [16]byte          // group domain
	UserID() [16]byte           // user domain
	Status() GroupInvite_Status //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type GroupInvite_Status uint8

const (
	GroupInvite_Status_Unset GroupInvite_Status = iota
	GroupInvite_Status_Active
	GroupInvite_Status_Inactive
)

// TODO::: user send last time of active state record plus its ID and optional invited user id as invite code.
