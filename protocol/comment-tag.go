

package protocol

import (
	"time"

)

// CommentTag indicate the domain record data fields.
// It must just use to index hashtag words from comment-text domain
type CommentTag interface {
	Tag() string         //
	CommentID() [16]byte // comment domain
	Time() protocol.Time // Save time
}

type CommentTag_StorageServices interface {
	Save(ct CommentTag) protocol.Error

	Count(tag string) (numbers uint64, err protocol.Error)
	Get(tag string, versionOffset uint64) (ct CommentTag, err protocol.Error)
	Last(tag string) (ct CommentTag, numbers uint64, err protocol.Error)
}
