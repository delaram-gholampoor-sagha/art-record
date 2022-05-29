

package protocol

import (
	"time"

)

// CommentText indicate the comment-text domain record data fields.
type CommentText interface {
	CommentID() [16]byte // comment domain
	Text() string        // Text With Style. HTML & CSS (No JS) is more expressive than markdown, suggest use them in article text to style text.
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}

type CommentText_StorageServices interface {
	Save(c CommentText) error

	Count(commentID [16]byte) (length uint64, err error)
	Get(commentID [16]byte, versionOffset uint64) (c CommentText, err error)
	Last(commentID [16]byte) (c CommentText, length uint64, err error)
}
