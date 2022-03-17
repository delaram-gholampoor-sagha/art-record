package protocol

type Music_Content struct {
	ID                  uint64 `json:"id"`
	Title               string `json:"title"`
	Acossiated_Category string `json:"acossiated_Category"`
	AlbumID             uint64 `json:"album_id"`
	Tag                 string `json:"tag"`
	Music_URL           string `json:"music_url"`
	Pic_URL             string `json:"pic_url"`
}

// can we filter our search ?

type Music interface {
	DeleteMusic(musicID uint64) error
	RegisterMusic(music Music) (Music, error)
	GetMusic(musicID uint64) (Music, error)
	GetMusics() ([]Music, error)

	SearchMusicByTag(tag string) ([]Music_Content, error)
	SearchMusicByCategory(Category string) ([]Music_Content, error)
	SearchMusicBySinger(singer string) ([]Music_Content, error)
	SearchMusicByTitle(title string) ([]Music_Content, error)
}

// approved by someone called admin ?? => maybe another microservice ?

