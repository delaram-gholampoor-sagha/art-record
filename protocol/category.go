package protocol

import "time"

type Category struct {
	ID uint64
	// ParentID   uint64
	// QuiddityID [16]byte
	// PictureID  uint64
	RelatedID      [16]byte
	Status         uint8
	Created_at     time.Time
	Updated_At     time.Time
}

type CategoryServices interface {
	CreateCategory(category Category) (Category, error)
	UpdateCategory(isParual bool, category Category) (Category, error)
	DeleteCategory(categoryID uint64) error
	GetCategory(CategoryID uint64) (Category, error)
	FindCategoryByTitle(title string) ([]Category, error)
}
