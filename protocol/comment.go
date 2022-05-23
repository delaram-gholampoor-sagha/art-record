

package protocol

import (
	"time"
)

// Comment indicate the comment domain record data fields.
// each comment is immutable record and so use version mechanism to chain comments in a group.
type Comment interface {
	CommentGroupID() [16]byte   // comment-group domain.
	CommentID() [16]byte        //
	ReplyTo() [16]byte          // other CommentID. Comment can be reply to other comment in the same group.
	UserID() [16]byte           // user-status domain
	Type() Comment_Type         //
	Settings() Comment_Settings //
	Time() time.Time            // Save time
	RequestID() [16]byte        // user-request domain
}

type Comment_StorageServices interface {
	Save(c Comment) error

	Count(commentGroupID [16]byte) (length uint64, err error)
	Get(commentGroupID [16]byte, versionOffset uint64) (c Comment, err error)
	Last(commentGroupID [16]byte) (c Comment, length uint64, err error)

	FindByReplyTo(commentGroupID [16]byte, commentID [16]byte, offset, limit uint64) (versionOffsets []uint64, length uint64, err error)
	FindByUserID(commentGroupID [16]byte, userID [16]byte, offset, limit uint64) (versionOffsets []uint64, length uint64, err error)

	ListUserGroups(userID [16]byte, offset, limit uint64) (ids [][16]byte, length uint64, err error)



	// protocol.EventTarget
}

type Comment_Type uint16

const (
	Comment_Type_Unset Comment_Type = 0
	Comment_Type_Text  Comment_Type = (1 << iota)
	Comment_Type_Voice // 0b00000010
	Comment_Type_Call  // 0b00000100
	Comment_Type_File  // 0b00001000

	Comment_Type_Image
	Comment_Type_Video
	Comment_Type_Album

	Comment_Type_Welcome
	Comment_Type_PhotoUpdated

	Comment_Type_Product

	Comment_Type_Thread
	Comment_Type_Forward // Retweet, Share, ...

	Comment_Type_NewConnection // made-friend, new-follower,
)

type Comment_Settings uint16

const (
	Comment_Settings_Unset        Comment_Settings = 0
	Comment_Settings_PreviewLinks Comment_Settings = (1 << iota) // Render web links as small widget or not
	Comment_Settings_Forwardable                                 // Allow to forward it
)
