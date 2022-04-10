package protocol

import "time"

type Comment struct {
	ID uint64
	//  This will determine which comment is for which page, it will be the page ID that you can specify on any webpage.
	// Page_ID uint64
	//  The parent comment ID that we'll use for comment replies.
	// Parent_ID   uint64
	// UserID      uint64
	// Post_ID     uint64
	// Quiddity_ID uint64
	RelatedID [16]byte

	CommentGroupID [16]byte // Use to retrieve comment. It can be URL hash or any random code.
	ReplyTo        [16]byte

	// The comment content will be what the user inputs via the form.
	Content string
	Status  uint8


	//  The date the comment was posted.
	Submit_Date time.Time
	Created_At  time.Time
	Updated_At  time.Time
}

// is it important to count how many times it has been published ?

type CommentServices interface {
	RegisterComment(Comment) (Comment, error)
	GetComments(commentgroupID uint64, offset uint64, limit uint64) ([]Comment, error)
	GetTheFirstComment(commentgroupID uint64) (Comment, error)
	UpdateComment(isPartial bool, comment Comment) (Comment, error)
	FindCommentByUserID(userID uint64) (Comment, error)
	FindTheParentComment(commentID uint64) (Comment, error)
	GetCommentState(commentID uint64) (string, error)
	DeleteComment(commentID uint64) error
}
