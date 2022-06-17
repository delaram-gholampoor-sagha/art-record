package protocol

// CommentText indicate the domain record data fields.
type CommentText interface {
	CommentID() [16]byte // comment domain
	Text() string        // Text With Style. HTML & CSS (No JS) is more expressive than markdown, suggest use them in article text to style text.
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type CommentText_StorageServices interface {
	Save(c CommentText) protocol.Error

	Count(commentID [16]byte) (numbers uint64, err protocol.Error)
	Get(commentID [16]byte, versionOffset uint64) (c CommentText, err protocol.Error)
	Last(commentID [16]byte) (c CommentText, numbers uint64, err protocol.Error)

	FindByText(text string, offset, limit uint64) (commentIDs [16]byte, numbers uint64, err protocol.Error)
}
