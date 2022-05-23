package protocol

import "time"

// CommentMentioned indicate the comment-mentioned domain record data fields.
type CommentMentioned interface {
	UserID() [16]byte    // user-status domain
	CommentID() [16]byte // comment domain
	Time() time.Time     // Save time
}

type CommentMentioned_StorageServices interface {
	Save(cm CommentGroup) error

	Count(mentionID [16]byte) (length uint64, err error)
	Get(mentionID [16]byte, versionOffset uint64) (cm CommentMentioned, err error)
	Last(mentionID [16]byte) (cm CommentMentioned, length uint64, err error)

	FindByUserID(userID [16]byte, offset, limit uint64) (ids [][16]byte, length uint64, err error)
}
