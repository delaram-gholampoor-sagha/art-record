package protocol

import "time"

type Category_Conent struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created_at time.Time `json:"created_at"`
}

type Category interface {
	CreateCategory(category Category_Conent) (Category_Conent, error)
	UpdateCategory(isParual bool , category Category) (Category, error)
	DeleteCategory(categoryID uint64) error
	GetCategory(CategoryID uint64) (Category_Conent, error)
	GetCategories() ([]Category_Conent, error)
}
