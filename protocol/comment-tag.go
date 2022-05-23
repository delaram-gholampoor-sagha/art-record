

package protocol

import (
	"time"

)

// CommentTag or hashtag indicate the comment-tag domain record data fields.
type CommentTag interface {
	Tag() string         //
	CommentID() [16]byte // comment domain
	Time() time.Time     // Save time
}

// do we even need this domain in our software ? as we don't tag in our software

type CommentTag_StorageServices interface {

}
