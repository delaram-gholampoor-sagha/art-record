package protocol

import "time"

type Order struct {
	ID uint64 `json:"id"`
	// The amount for which the order was created, in currency subunits. For example, for an amount of ₹295, enter 29500.
	Amount uint64 `json:"amount"`
	// boolean Indicates whether the customer can make a partial payment. Possible values:
	// true: The customer can make partial payments.
	// false (default): The customer cannot make partial payments.
	Partial_payment bool `json:"partial_payment"`
	// The amount paid against the order.
	Amount_paid uint64 `json:"amount_paid"`
	// The amount pending against the order.
	Amount_due uint64 `json:"amount_due"`
	// The currency associated with the order's amount. The default length is 3 characters. Refer to the list of supported currencies.
	Currency string `json:"currency"`
	// Receipt number that corresponds to this order. Can have a maximum length of 40 characters and has to be unique.
	Receipt string `json:"receipt"`
	// Possible values:
	// created: When you create an order it is in the created state. It stays in this state till a payment is attempted on it.
	// attempted: An order moves from created to attempted state when a payment is first attempted on it. It remains in the attempted state till one payment associated with that order is captured.
	// paid: After the successful capture of the payment, the order moves to the paid state.
	// No further payment requests are permitted once the order moves to the paid state.
	// The order stays in the paid state even if the payment associated with the order is refunded.
	Status   string `json:"status"`
	Attempts uint64 `json:"attempts"`
	//The number of payment attempts, successful and failed, that have been made against this order.
	// json object Key-value pair that can be used to store additional information about the entity.
	// Maximum 15 key-value pairs, 256 characters (maximum) each. For example, "note_key": "Beam me up Scotty”.
	Notes      string    `json:"notes"`
	// Returns a collection of all payments made for each order.
	// Payments    
	// Returns the card details of each payment made for each order.
	// Payments_Cards  
	// Returns a collection of transfers created for each order.
	// Transfers   
	// Returns the virtual account details created for each order.
	// Virtual_Account
	Created_At time.Time `json:"created_at"`
}






// edit order (check out page )
//    Check all your order details before continuing to checkout & pay.
// specific order details
//    Provide the best experience, allow them to set up time for delivery, payment method & driver tip.

// Pre-order functionality
// Allow your customers to order in advance to any of your business, all they need to do is select the date & time of delivery,
// you get the order in your Orders Manager and that's it, are you ready to get a lot of orders?

// =============================================================================

// About Orders
// Order is an important step of the payment life cycle at Razorpay. When a customer clicks the pay button on your website or app,
//an order is created with a unique identifier. This contains details such as the transaction amount and currency. Pass this order ID to the Razorpay Checkout.

// You need to integrate your server with Orders API before proceeding with Razorpay Payment Gateway integration on your website or app.

// Advantages
// Single successful payment bound to an order. Prevents multiple payments.
// Quick and easy query in the database. Combines multiple payment attempts for a single order.
// Order States
// Following are the various states of an order:




