

package protocol

import (
	"time"

)

// CommentStatus indicate the comment-status domain record data fields.
type CommentStatus interface {
	CommentID() [16]byte          // comment domain
	Status() CommentStatus_Status //
	Time() time.Time              // Save time
	RequestID() [16]byte          // user-request domain
}

type CommentStatus_StorageServices interface {
	Save(cs CommentStatus) error

	Count(commentID [16]byte) (length uint64, err error)
	Get(commentID [16]byte, versionOffset uint64) (cs CommentStatus, err error)
	Last(commentID [16]byte) (cs CommentStatus, length uint64, err error)


	// protocol.EventTarget
}

type CommentStatus_Status uint8

const (
	CommentStatus_Status_Unset   CommentStatus_Status = 0
	CommentStatus_Status_Created CommentStatus_Status = (1 << iota)
	CommentStatus_Status_Edited
	CommentStatus_Status_Locked
	CommentStatus_Status_Hidden
	CommentStatus_Status_Deleted
	CommentStatus_Status_Blocked
)
