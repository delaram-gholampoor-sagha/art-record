package protocol

type Album_Content struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Pic_URL string `json:"pic_url"`
}

type Album interface {
	CreateAlbum(album Album_Content) (Album_Content, error)
	DeleteAlbum(albumID uint64) error
	GetAlbum(albumID uint64) (Album_Content, error)
	UpdateAlbum(isPartial bool, albumID uint64)
	FindAlbumByTitle(name string) (Album_Content, error)
}
