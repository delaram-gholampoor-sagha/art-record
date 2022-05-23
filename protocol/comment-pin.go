

package protocol

import "time"

// CommentPin indicate the comment-pin domain record data fields.
type CommentPin interface {
	GroupID() [16]byte   // comment-group domain.
	CommentID() [16]byte // comment domain
	Status() CommnetPin_Status
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}


type CommentPin_StorageService interface {
	Save(cp CommentPin) error

	Count(commentID [16]byte) (length uint64, err error)
	Get(GrouopID [16]byte, versionOffset uint64) (groupids [16]byte, err error)
	Last(commentID [16]byte) (cp CommentPin, length uint64, err error)

   

	FindByGroupID(groupIDs  [][16]byte) (id [16]byte , err error)
}

type CommnetPin_Status uint8

const (
    CommnetPin_Status_Unset = iota
	CommnetPin_Status_Unpin
	CommnetPin_Status_Pin
)