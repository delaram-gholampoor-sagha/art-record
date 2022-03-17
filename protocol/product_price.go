package protocol

type Product_Price_Content struct {
	ID    uint64 `json:"id"`


	Price uint64 `json:"price"`
	// not sure ...
  Discounted_Price string `json:"discounted_price"`
	SoldPrice string `json:"sold_price"`

}

type Product_Price interface {
	GetProductPrice(productID uint64) (Product_Price_Content, error)
	// isPartial ??
	UpdatePrice(productID uint64 ) (Product_Price_Content, error)
	DeletePrice(productID uint64 ) error 
}


// actual cost = > بهای تمام شده -- هزینه ی واقعی
// average cost = >  هزینه ی میانگین
// landed cost => هزینه ی تمام شده در انبار
// manifacturing cost => هزینه ی ساخت
// marginal cost => هزینه یابی نهایی
// overall cost => هزینه ی کل
// prime the pump => پول توی کاری ریختن
// response cost => جریمه
// standard cost => هزینه ی معیار
// unit cost => بهای واحد
// unit labor cost => هزینه ی واحد کار
// weighted average cost => هزینه ی میانگین موزون


