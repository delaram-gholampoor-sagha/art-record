package protocol

import "time"

type Post struct {
	ID         [16]byte
	QuiddityID [16]byte
	Status     uint8
	Created_At time.Time
	Updated_At time.Time
}



// It depends on the context
// State generally refers to the entire state of an entity
// - all its values and relationships at a particular point in time (usually, current)
// Status is more of a time-point, say, where something is at in a process or workflow
// - is it dirty (therefore requiring saving), is it complete, is it pending input, etc

type PostServices interface {
	GetPost(postID uint64) (post Post, err error)
	CreatePost(post Post) (Post, error)
	// the last 3 one
	GetPosts() ([]Post, error)
	GetPostStatus(postId uint64) (string, error)
	FindPostByTitle(title string) (Post, error)
	DeletePost(postID int) error
	UpdatePost(isPartial bool, post Post) (Post, error)
}
