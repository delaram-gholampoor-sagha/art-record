package protocol

import "time"

// it is cancellable ?
type Invoice_Cart struct {
	ID             uint64
	UserID         uint64
	Status         uint8
	CartType       uint8
	Line_Item      Invoice_Line_Item
	// Actual_Cost    uint64
	Total_Amount   uint64
	// the number of times that the user tries to add sth to the cart
	// or try to create the cart ?!!
	Attempts       uint64
	Created_at     time.Time
	Updated_at     time.Time
	// do we want to give this ability to the user ?
	Cancelled_At   time.Time
	// do we want to expire this at any point in time ?
	// or it could be cancellable just by user's authority ?
	ExpirationDate time.Time
}

type CartServices interface {
	GetCartState(cartID uint64) (string, error)
	UpdateCart(isPartial bool, cart Invoice_Cart) (Invoice_Cart, error)
	DeleteCart(cartID uint64) (Invoice_Cart, error)
	GetCarts() ([]Invoice_Cart, error)
	FindCartByType(cartType string) (Invoice_Cart, error)
	PricePerUnit(unit string) (string, error)
	RegisterCart(cart Invoice_Cart) (Invoice_Cart, error)
	AddItemsToCart()
	RemoveAllItemsFromCart(number uint64) (uint64, error)
	IncreaseItem(number uint64) (uint64, error)
	DecreaseItem(number uint64) (uint64, error)
	// product inventory ?
	CheckoutAvailability()
}
