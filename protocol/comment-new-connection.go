package protocol

// CommentFollow indicate the domain record data fields.
type CommentFollow interface {
	CommentID() [16]byte      // comment domain
	FollowedUserID() [16]byte // user domain. FollowedUserID
	Type() CommentFollow_Type //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type CommentFollow_Type uint8

// CommentFollow_Type_MakeFriend
// CommentFollow_Type_Follow
