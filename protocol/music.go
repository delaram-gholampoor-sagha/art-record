package protocol

import "time"

// isnt it the same as audio ?
type Music struct {
	ID                     uint64
	Quiddity_ID            [16]byte
	// ???
	Song_Related_People_ID [16]byte
	// Category_ID    uint64
	// Album_ID       uint64
	// Tag_ID         uint64
	// PictureID      uint64
	// AudioID        uint64
	// Song_Related_People  uint64
	Status     uint8
	Created_At time.Time
	Updated_At time.Time
}

type MusicService interface {
	DeleteMusic(musicID uint64) error
	RegisterMusic(music Music) (Music, error)
	UpdateMusic(isPartial bool, music Music) (Music, error)

	SearchMusicByTag(tag string) ([]Music, error)
	SearchMusicByCategory(Category string) ([]Music, error)
	SearchMusicBySinger(singer string) ([]Music, error)
	SearchMusicByAlbum(title string) ([]Music, error)
}
