package protocol


type Product_Price struct {
	ID               [16]byte
	RelatedID        [16]byte
	Prime_Price      int64
	// Sold_Price       uint64
	Status           uint8
}





type Product_Price_Services interface {
      GetDiscountedPrice()
}
