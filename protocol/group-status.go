package protocol

// GroupStatus indicate the domain record data fields.
type GroupStatus interface {
	GroupID() [16]byte    // group-status domain
	Status() Group_Status //
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type GroupStatus_StorageServices interface {
	Save(gs GroupStatus) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (gs GroupStatus, err protocol.Error)
	Last(groupID [16]byte) (gs GroupStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status Picture_Status, offset, limit uint64) (groupIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}

type Group_Status Quiddity_Status

const (
// Group_Status_ = Group_Status(Quiddity_Status_FreeFlag << iota)
)
