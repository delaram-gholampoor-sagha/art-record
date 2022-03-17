package protocol

import "time"

type Payment struct {
	ID     uint64 `json:"id"`
	Entity string `json:"entity"`
	// The payment amount represented in smallest unit of the currency passed.
	// For example, amount = 100 translates to 100 paise, that is ₹1 (default currency is INR).
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
	// Possible values:
	// -created
	// -authorized
	// -captured
	// -refunded
	// -failed
	Status string `json:"status"`
	// The converted payment amount used to calculate fees and settlements. Represented in smallest unit of the base_currency.
	Base_Amount uint64 `json:"base_amount"`
	// The conversion currency used to calculate fees and settlements
	Base_Currency uint64 `json:"base_currency"`
	// The payment method used for making the payment.
	// Possible values:
	// -card
	// -netbanking
	// -wallet
	// -emi
	// -upi
	Method      string `json:"method"`
	Order_Id    uint64 `json:"order_id"`
	Description string `json:"description"`
	// Indicates whether the payment is done via an international card or a domestic one.
	International bool `json:"international"`
	// Possible values include:
	// - null
	// - partial
	// - full
	Refund_Status string `json:"refund_status"`
	// The amount refunded in smallest unit of the currency passed.
	// For example, if amount_refunded = 100, here 100 stands for 100 paise, which is equal to ₹1. INR is the default currency.
	Amount_Refunded uint64 `json:"amount_refunded"`
	// Indicates if the payment is captured.
	Captured string `json:"captured"`
	// Customer email address used for the payment
	Email string `json:"email"`
	// Customer contact number used for the payment.
	Contact string `json:"contact"`
	//  Fee (including GST) charged by Razorpay.
	Fee uint64 `json:"fee"`
	// GST charged for the payment.
	Tax uint64 `json:"tax"`
	// Error that occurred during payment. For example, BAD_REQUEST_ERROR.
	Error_Code string `json:"error_code"`
	// Description of the error that occurred during payment. For example, Payment processing failed because of incorrect OTP.
	Error_Description string `json:"error_description"`
	// The point of failure. For example, customer.
	Error_Source string `json:"error_source"`
	// The stage where the transaction failure occurred.
	// The stages can vary depending on the payment method used to complete the transaction.
	// For example, payment_authentication.
	Error_Step string `json:"error_step"`
	// The exact error reason. For example, incorrect_otp.
	Error_Reason string `json:"error_reason"`
	// Contains user-defined fields, stored for reference purposes.
	Notes      string    `json:"notes"`
	Created_At time.Time `json:"created_at"`
	// Expanded Offer details, applicable when an offer was applied to the payment.
	Offers string `json:"offers"`
	// Expanded card details, usable for card and EMI payments.
	Card string `json:"card"`
	// Expanded EMI plan details, usable for EMI payments.
	// An equated monthly instalment (EMI) is a set monthly payment provided by a borrower to a creditor on a set day, each month.
	// EMIs apply to both interest and principal each month, and the loan is paid off in full over some years
	// اقساط ماهانه معادل (EMI) پرداخت ماهانه مشخصی است که توسط وام گیرنده در یک روز معین و هر ماه به یک بستانکار ارائه می شود.
	// EMI هر ماه برای هر دو سود و اصل اعمال می شود و وام به طور کامل طی چند سال پرداخت می شود
	Emi string `json:"emi"`
}

// Capture a Payment
// After the customer's bank authorises the payment,
//you must verify if the authorised amount deducted from the customer's account
// is the same as the amount paid by the customer on your website or app.

// Fetch Payments Based on Orders
// /orders/:id/payments

// Fetch Card Details of a Payment
// /payments/:id/card

// Path Parameter
// id : string Unique identifier of the payment for which you want to retrieve card details.
// mandatory
// 
// Response Entities

// id : string Unique identifier of the card used for the payment.
// 
// entity : string The value for this attribute will always be card.

// name :  string Name of the cardholder.

// last4 : string The last 4 digits of the card number.

// network : 
//   Possible values:
//   American Express
//   Diners Club
//   Maestro
//   MasterCard
//   RuPay
//   Unknown
//   Visa
// type : The card type.
//    Possible values:
//    credit
//    debit
//    prepaid
//    unknown
// issuer : The card issuer - The 4-character code denotes the issuing bank
// international : boolean This attribute will be set to true if the card is issued by a foreign bank.
// emi : boolean This attribute is set to true if the card can be used for EMI payment method.
// sub_type : boolean The sub-type of the customer's card.
//   Possible values:
//      business
//      Corporate Card : ? =>  // Corporate credit cards are credit cards issued to employees of established companies that let them 
// charge their authorized business expenses—such as hotel stays and plane tickets—without having to use their own card or cash.
//       consumer



