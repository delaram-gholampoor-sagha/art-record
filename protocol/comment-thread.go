

package protocol

import (
	"time"

)

// CommentThread indicate the comment-thread domain record data fields.
type CommentThread interface {
	CommentID() [16]byte             // comment domain
	Name() string                    //
	ArchiveAfter() protocol.Duration //
	Time() time.Time                 // Save time
	RequestID() [16]byte             // user-request domain
}


type CommentThread_StorageServices interface {
	Save(c CommentText) error

	Count(commentID [16]byte) (length uint64, err error)
	Get(commentID [16]byte, versionOffset uint64) (c CommentText, err error)
	Last(commentID [16]byte) (c CommentText, length uint64, err error)
}


