package protocol

// Do we ever need tags in this software ?
type Tag struct {
	ID      uint64
	Content string
	Status  string
}

type TagServices interface {
	RegisterTag(tag Tag) (Tag, error)
	UpdateTag(rag string) (Tag, error)
	DeleteTag(tagID uint64) error
	FindTagByTitle(content string) ([]Tag, error)
}
