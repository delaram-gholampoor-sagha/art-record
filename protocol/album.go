package protocol

type Album_Content struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Album interface {
	CreateAlbum(album Album_Content) (Album_Content, error)
	DeleteAlbum(albumID uint64) error
	SearchAlbum(title string) (Album_Content, error)
	GetAlbums() ([]Album_Content, error)
	GetAlbum(albumID uint64) (Album_Content, error)
}
