package protocol

import "time"

type Post struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// Is it published or not ?
	Status     bool      `json:"status"`
	Pic_URL    string    `json:"pic_url"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

// I'm thinking about a draft state where the admin could think about the content
// and maybe publish the post later on ...

type Blog interface {
	// blog
	GetPost(postID uint64) (post Post, err error)
	CreatePost(post Post) (Post, error)
	GettheLastThreePosts() ([]Post, error)
	GetPostStatus(postId uint64) (string, error)
	FindPostByTitle(title string) (Post, error)
	DeletePost(postID int) error
	UpdatePost(isPartial bool, post Post) (Post, error)
}
