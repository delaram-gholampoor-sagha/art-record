package protocol

type Product_Price_Content struct {
	ID    uint64 `json:"id"`
	Price uint64 `json:"price"`
	// not sure ...
	Discounted_Price string `json:"discounted_price"`
}

type Product_Price interface {
	GetProductPrice(productID uint64) (Product_Price_Content, error)
	// isPartial ??
	UpdatePrice(productID uint64 ) (Product_Price_Content, error)
	DeletePrice(productID uint64 ) error
}
