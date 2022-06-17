package protocol

// CommentStatus indicate the domain record data fields.
type CommentStatus interface {
	CommentID() [16]byte    // comment domain
	Status() Comment_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type CommentStatus_StorageServices interface {
	Save(cs CommentStatus) protocol.Error

	Count(commentID [16]byte) (numbers uint64, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (cs CommentStatus, err protocol.Error)
	Last(commentID [16]byte) (cs CommentStatus, numbers uint64, err protocol.Error)

	// FilterByStatus(status Comment_Status, offset, limit uint64) (commentIDs [][16]byte, numbers uint64, err protocol.Error)
	protocol.EventTarget
}

type Comment_Status Quiddity_Status

const (
// Comment_Status_ = Comment_Status(Quiddity_Status_FreeFlag << iota)
)
