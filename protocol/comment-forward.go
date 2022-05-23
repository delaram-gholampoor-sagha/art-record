package protocol

import "time"

// CommentForward indicate the comment-forward domain record data fields.
type CommentForward interface {
	CommentID() [16]byte // comment domain
	Forwarded() [16]byte // comment domain. or ForwardedCommentID
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}

type CommentForward_StorageServices interface {
	Save(c CommentForward) (err error)

	Count(commentID [16]byte) (length uint64, err error)
	Get(commentID [16]byte, versionOffset uint64) (ci CommentForward, err error)
	Last(commentID [16]byte) (ci CategoryItem, length uint64, err error)

	FindByForwardedID()
}
