package protocol

type Album struct {
	ID         [16]byte
	QuiddityID [16]byte
	Status     uint8
}

type AlbumServices interface {
	CreateAlbum(album Album) error
	DeleteAlbum(albumID uint64) error
	GetAlbum(albumID uint64) (Album, error)
	UpdateAlbum(isPartial bool, albumID uint64)
	FindAlbumByTitle(name string) (Album, error)
}
