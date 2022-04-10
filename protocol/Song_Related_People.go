package protocol

type Song_Related_People struct {
	ID uint64
	// AlbumID    uint64
	// MusicID    uint64
	// QuiddityID uint64
	RelatedID  [16]byte
	// ترانه‌سرا
	SongWriter string
	// آهنگ‌ساز
	Composer   string
	// تنظیم کننده اهنگ
	Regulator  string
	// کارگردان
	Director     string
	// تهیه‌کننده موسیقی
	MusicSupplier  string
	Status  uint8
}

type Song_Related_People_Services interface {
	RegisterSinger()
	DeleteSinger()
	UpdateSinger()
	// FindSingerByAlbum()
	// FindSingerByMusic()
}
