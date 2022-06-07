package protocol

// CommentGroup indicate the domain record data fields.
type CommentGroup interface {
	GroupID() [16]byte                  // or ChannelID. Use to store and retrieve comments in time order. It can be UUID, URL hash or any random id.
	OwnerUserID() [16]byte              // user-status domain
	Time() protocol.Time                // Save time
	RequestID() [16]byte                // user-request domain
}

type CommentGroup_StorageServices interface {
	Save(cg CommentGroup) protocol.Error

	Count(groupID [16]byte) (numbers uint64, err protocol.Error)
	Get(groupID [16]byte, versionOffset uint64) (cg CommentGroup, err protocol.Error)
	Last(groupID [16]byte) (cg CommentGroup, numbers uint64, err protocol.Error)

	FindByUserID(ownerUserID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err protocol.Error)
}
