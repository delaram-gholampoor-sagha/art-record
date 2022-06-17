package protocol

// GroupOwner indicate the domain record data fields.
type GroupOwner interface {
	GroupID() [16]byte     // or ChannelID. Use to store and retrieve comments in time order. It can be UUID, URL hash or any random id.
	OwnerUserID() [16]byte // user domain
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}

type GroupOwner_StorageServices interface {
	Save(g GroupOwner) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (g GroupOwner, err protocol.Error)
	Last(groupID [16]byte) (g GroupOwner, numbers uint64, err protocol.Error)

	FindByUserID(ownerUserID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
