package protocol

import "time"

type Post struct {
	ID         uint64
	Title      string
	Content    string
	// is it published or not ?
	Status     bool
	Pic_URL    string
	Created_At time.Time
	Updated_At time.Time
}

type Blog interface {
	// blog
	GetPost(postID uint64) (post Post, err error)
	CreatePost(post Post) (Post, error)
	GetPosts() (posts []Post, err error)
	SearchPost(title string) ([]Post, error)
	DeletePost(postID int) error
	UpdatePost(isPartial bool, post Post) (Post, error)
}
