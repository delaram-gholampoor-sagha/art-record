package protocol

import "time"

type Cart_Content struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	// ProductID       string `json:"product_id"`
	Status         string `json:"state"`
	Unit           string `json:"Unit"`
	Payable_Amount string `json:"payable_amount"`
	// Sum_Total      string `json:"sum_total"`
	Created_at     string    `json:"created_at"`
	Updated_at     string    `json:"updated_at"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type Cart interface {
	GetCartState(cartID uint64) (string, error)
	UpdateCart(isPartial bool, cart_content Cart_Content) (Cart_Content, error)
	DeleteCart(cartID uint64) (Cart_Content, error)
	// different type of carts ?
	GetCarts() ([]Cart_Content, error)
	FindCartByType(cartType string) (Cart_Content, error)
	PricePerUnit(unit string) (string, error)
	RegisterCart(cart Cart_Content) (Cart_Content, error)
	AddItemsToCart()
	RemoveItemsFromCart(number uint64) (uint64, error)
	IncreaseItem(number uint64) (uint64, error)
	DecreaseItem(number uint64) (uint64, error)
	CheckoutAvailability()
}

// سبد خرید شکل گرقت
// چند تا ایتم داخلش ریخته شد
// کاربر اومد یه ایتم رو حذف کرد
// تصیم گرفت خرید رو نهایی کند
// میخاد به صورت انلاین پرداخت کند
// کاربر موقع خرید صفحه را رفرش کرد و ما را بدبخت کرد
// کاربر ما را سره کار گذاشته و نمی خاهد سبد خرید خود را نهایی کند ولش میکند به امان خدا

// inventory ?

// Below is the list of features offered by an advanced shopping cart:

// Products page information like customer reviews and ratings
// View cart page for viewing the selected products
// Search functionality
// The integrated blog features to increase your readers
// Backend features offer by an advanced shopping cart:

// Invoice management
// Inventory management
// Order management
// Real-time rate calculation and Shipping integration
// Tax calculation
// Marketplace integrations
// Social media integrations such as Google Shopping, Facebook, Instagram
// Variety of payment options

// type Item struct {
// 	Name           string
// 	Units          int
// 	PricePerUnit   float32
// 	ExpirationDate time.Time
//  }

// type ItemTotal struct {
// 	*cart.Item
// 	total float32
//  }

// Filter: We only need to leave in the cart those items whose expiration date has not expired yet.
// Sort: We need to print the items ordered by name
// Map: For any item we need to add a new attribute, the total price of the items
// ForEach: For each item in the cart we need to print the required info
// Reduce: We need to calculate the total price for our shopping cart.

//Product for holding a single product details
// type Product struct {
// 	ID          string   `bson:"_id,omitempty"`
// 	Code        string   `bson:"code"`
// 	Name        string   `bson:"name"`
// 	Description string   `bson:"description"`
// 	Price       float64  `bson:"price"`
// 	Count       int      `bson:"count"`
// 	Discount    float64  `bson:"discount"`
// 	Colors      []string `bson:"colors"`
// 	Sizes       []string `bson:"sizes"`
// }

// //Item for holding a single item details
// type Item struct {
// 	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Code  string             `json:"code" bson:"code"`
// 	Unit  int                `json:"unit" bson:"unit"`
// 	Price float64            `json:"price" bson:"price"`
// 	Total float64            `json:"total" bson:"total"`
// }

// //Order for holding a single order details
// type Order struct {
// 	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Items []Item             `json:"items" bson:"items"`
// 	Total float64            `json:"total" bson:"total"`
// }

// //ResponseData for holding a api response details
// type ResponseData struct {
// 	Status  string `json:"status"`
// 	ID      string `json:"id,omitempty"`
// 	Message string `json:"message"`
// }
