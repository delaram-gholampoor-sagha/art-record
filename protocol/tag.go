package protocol

type Tag_Content struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
}

type Tag interface {
	RegisterTag(tag Tag_Content) (Tag_Content, error)
	UpdateTag(rag string) (Tag_Content , error)
	DeleteTag(tagID uint64) error
	SearchTags(content string ) ([]Tag_Content, error)
}
