package protocol

import "time"

type Settlement struct {
	ID uint64 `json:"id"`
	// The name of the entity
	Entity string `json:"entity"`
	// The settlement amount in paise.
	Amount uint64 `json:"amount"`
	// there are some conflicts that needs to be checked.
	// Possible values:
	// created
	// processed
	// failed
	Status string `json:"status"`
	// This is the total fees charged to process all payments received from customers settled to you in this settlement transaction.
	Fees uint64 `json:"fees"`
	// The total tax, in paise, charged on the fees collected to process all payments 
	// received from customers settled to you in this settlement transaction.
	Tax uint64 `json:"tax"`
	// Unique Transaction Reference number available across banks,
	// which can be used to track a particular settlement in your bank account. For example, 1597813219e1pq6w.
	UTR        string    `json:"utr"`
	Created_At time.Time `json:"created_at"`
}

// About Settlements
// Settlement is the process in which the money received from your customers is settled in your bank account

// Domestic Settlements
// The standard settlement cycle for domestic payments is T+2 working days (where T is the date of transaction capture).

// Example
// You captured 20 payments on February 02, 2019 at 9:00 a.m. and your settlement schedule is T+2 days.
// The payments you captured on February 02, 2019 will be settled to your bank account on February 04, 2019 at 9:00 a.m.
// If the settlement day is a bank holiday, the settlement will be made on the next working day after the bank holiday.

// International Settlements
// The standard settlement cycle for international payments is T+7 working days (where T is the transaction capture date).

// Example
// You captured 7 international payments on February 02, 2019 and your settlement schedule is T+7 days 9:00 a.m.
// The payments you captured on February 02, 2019 will be settled to your account on February 09, 2019 at 9:00 a.m.
// If the settlement day is a bank holiday, the settlement will be made on the next working day after the bank holiday.

// Instant Settlements
// Using Instant Settlements, you can get access to your funds as and when you want.
// Normally, you would receive settlements according to your settlement cycle.

// Settlement States
// Following are the various states of a settlement:
// processed - failed

// Some reasons for settlement failure:
// Your bank account is inactive or frozen.
// Incorrect bank account details.
// The settlement was rejected by your bank.

// =====================================================================================
// Settlements APIs
// By default, the complete process takes a time of T+2 business days for domestic transactions and T+7 days for international transactions,
// T being the date of capture of payment.

// Settlement Recon
// The settlement recon API returns a list of all transactions such as payments, refunds, transfers
// and adjustments that have been settled to you for a particular day or month.
// /settlements/recon/combined?year=yyyy &month=mm


// type => string |: Indicates the type of transaction. Possible values: 
// payment
// refund
// transfer
// adjustment
// debit

// on_hold =>  boolean /: Indicates whether the account settlement for transfer is on hold.
// Possible values:
// true
// false


// settled => boolean Indicates whether the transaction has been settled or not.
// Possible values:
// true
// false
// created_at
// integer Timestamp (in Unix) at which the transaction was created.

// settled_at
// integer Timestamp (in Unix) when the transaction was settled.


// settlement_utr
// string Unique reference number linked to the settlement. For example, KKBKH14156891582.


// order_receipt
// string Receipt number entered while creating the Order.

// method
// string The payment method used to complete the payment.
// Possible values:
// card
// netbanking
// wallet
// emi
// upi

// card_network =>  string |: The card network used to process the payment. Possible values are:
// American Express
// Diners Club
// Maestro
// MasterCard
// RuPay
// Visa
// unknown

// card_issuer => string |: The card issuer. 4-character code denoting the issuing bank. For example, KARB.
// This attribute will not be set for international cards, that is, for cards issued by foreign banks.

// card_type
// string The card type used to process the payment. Possible values:
// credit
// debit

// dispute_id => string |: Unique identifier of any dispute that might have arisen for this transaction.
