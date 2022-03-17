package protocol

import "time"

type Route struct {
	ID     uint64 `json:"id"`
	Entity string `json:"entity"`
	// Transfer Status
	// The following table lists transfer statuses and their description:
	// Status : Description
	// created : As soon as the transfer is initiated, it moves to the created state.
	// Transfers via payments and direct transfers will immediately move to the pending state from the created state.
	// However, the transfer will stay in the created state for transfers via orders until the order is paid.
	// pending : The transfer will be in the pending state till the transfer is waiting to be processed.
	// processed : Once the transfer is processed, the status will change from pending to processed.
	// You can initiate the reversal if required.
	// failed : If the transfer processing fails because of errors like insufficient balance, the transfer status will change to the failed state.
	// You can proceed to create a new transfer request.
	// reversed : If the complete reversal is done for the transfer, then it will be marked as reversed.
	// partially_reversed : If the partial reversal is done on transfer, it will be marked as partially_reversed.
	Transfer_Status string `json:"transfer_status"`
	// Settlement Status
	// The following table lists settlement statuses and their description:
	// Status : Description
	// pending : This status indicates that the transfer is processed and the settlement is pending.
	// You can create the reversal if required.
	// on_hold : If you have used the on_hold parameter to hold settlements for the transfer
	// and don’t pass the on_hold_until parameter to define till when it should be on_hold, the transfer will be on hold indefinitely.
	// We will mark such transfers as on_hold.
	// settled : When the transfer is settled to the Linked Account.
	Settlement_status string `json:"settlement_status"`
	// Unique identifier of the transfer source. The source can be a payment or an order.
	Source string `json:"source"`
	// Unique identifier of the transfer destination, that is, the linked account.
	Recipient string `json:"recipient"`
	// The amount to be transferred to the linked account, in paise.
	Amount uint64 `json:"amount"`
	// ISO currency code.
	Currency string `json:"currency"`
	// Amount reversed from this transfer for refunds
	Amount_Reveresed uint64 `json:"amount_reveresed"`
	// // json object Set of key-value pairs that can be associated with an entity.
	// These pairs can be useful for storing additional information about the entity.
	// A maximum of 15 key-value pairs, each of 256 characters (maximum), are supported.
	// For example, "region": "south", "city": "Bangalore".
	Notes string `json:"notes"`
	// Provides error details that may occur during transfers
	Error Error `json:"error"`
	// array List of keys from the notes object which needs to be shown to linked accounts on their Dashboard.
	// For example, "region", "city". Only the keys will be shown, not values.
	Linked_Account_Notes string `json:"linked_account_notes"`
	// Indicates whether the account settlement for transfer is on hold.
	// Possible values:
	// 1 - Puts the settlement on hold.
	// 0 - Releases the settlement.
	On_Hold bool `json:"on_hold"`
	// Timestamp, in Unix format, indicates until when the settlement of the transfer must be put on hold.
	// If no value is passed, the settlement is put on hold indefinitely.
	On_Hold_Until uint64 `json:"on_hold_untill"`
	// Unique identifier of the settlement.
	Recipient_SettleMent_ID uint64    `json:"recipient"`
	Created_At              time.Time `json:"created_at"`
}

// =======================================================

// Provides error details that may occur during transfers
type Error struct {
	ID uint64 `json:"id"`
	// Type of the error.
	Code string `json:"code"`
	// Error description.
	Description string `json:"description"`
	// Name of the parameter in the API request that caused the error.
	Field string `json:"field"`
	// The point of failure in the specific operation. For example, customer, business and so on.
	Source string `json:"source"`
	// The stage where the transaction failure occurred.
	// Stages can be different depending on the payment method used to make the transaction.
	Step string `json:"step"`
	// The exact error reason. It can be handled programmatically
	Reason string `json:"reason"`
}

// Route transfers can be performed only on captured payments.

// Create Transfers from Orders
// POST  orders
// REQUEST PARAMETERS ??
// amount : The transaction amount, in paise. For example, for an amount of ₹299.35, the value of this field should be 29935.
// mandatory
// currency : The currency in which the transaction should be made. We support only INR for Route transactions.
// mandatory
// receipt : Unique identifier that you can use for internal reference.
// optional ??
// transfers : Details regrading to transfer

// Create Transfers from Payments
// 1_The customer pays the amount via normal payment flow.
// 2_Once the payment is captured, you can initiate a transfer to linked accounts with a transfer API call.
// You have to specify the details of the account_id and amount
// POST   /payments/:id/transfers

// Direct Transfers
// you can also transfer funds to your linked accounts directly from your account balance using the Direct Transfers API.
// POST /transfers

// Fetch Transfers for a Payment
// Use this endpoint to fetch the collection of all transfers created on a specific Payment ID.
// GET /payments/:id/transfers?status

// Fetch Transfer for an Order
// Use this endpoint to fetch the collection of all transfers created on a specific Order ID.
// GET /orders/:id/?expand[]=transfers &status

// Fetch Transfers for a Settlement
// You can use the following endpoint to retrieve the collection of all transfers made for a particular recipient_settlement_id
// GET /transfers

// Fetch Settlement Details
// You can use the following endpoint to fetch the details of settlements made to linked accounts.
// You must append ?expand[]=recipient_settlement as the query parameter to the fetch transfer request.
// This would return a settlement entity along with the transfer entity.
// GET /transfers?expand[]=recipient_settlement

// Fetch Payments of a Linked Account
// You can use the following endpoint to fetch a list of all the payments received by a linked account.
// GET /payments/

// Refund Payments and Reverse Transfer from a Linked Account
// A payment can be refunded to a customer using Refunds API.

// When refunding a payment that has transfers, the amount is deducted from your main account balance. You can set the reverse_all parameter to 1 in the refund POST request to recover the amount from the linked account. This process will recover the amount from the linked account for every transfer made on the payment before processing the refund to the customer.

// Reversals can be automated with the reverse_all attribute in the following refund scenarios:

// Full refund
// Partial refund for a payment transferred to a single account.

// Reverse Transfers from all Linked Accounts
// While a transfer moves funds from your account to linked accounts, a reversal can move funds back into your account.
// You can use the following endpoint to create reversals on a particular transfer_id.
// POST /transfers/:id/reversals

// This request creates a reversal entity and reverses the funds transferred from your account to the linked account.
// Here, the amount specified is debited from the linked account balance and credited to your balance.
// Partial reversals are also supported, and you can create multiple reversals on a transfer_id.
// If you do not provide the amount parameter in the request, then the entire amount of the transfer is reversed.

// Linked Account Settlements
// You have complete control over settlements to linked accounts.
// Transfers are settled to linked accounts as per the settlement_period defined for the account.
// You can also choose to hold account settlements per transfer or per linked account.

// Hold Settlements For Transfers
// When transferring payment to an account, you can choose to put the transfer settlement on hold indefinitely or until a defined time.
// You can change these settings anytime via the provided API until the settlement is made.
// When you put a transfer settlement on hold, the settlement will not happen until you release it.
// You can use the following endpoint to create a transfer with settlement configurations:
// POST /payments/:id/transfers
// Settlement Schedule Override:
// The settlement schedule defined for the linked account takes precedence over the on_hold and on_hold_until functionality.
// This simply means that a defined settlement schedule is the minimum time required for the transfer to be settled.
// Examples: For a T+10 settlement schedule:

// If you create a transfer with on_hold: 1 and then release it on T+7 day, the settlement will only go out on T+10 day.
// If you create a transfer with on_hold: 1 and on_hold_until: 1491567400 (assume the timestamp 1491567400 corresponds to 7 days after transfer),
// then the on_hold will change to false on T+7 day, but the settlement will only go out on T+10 day.

// Modify Settlement Hold for Transfers
// You can use the following endpoint to modify the settlement configuration for a particular transfer_id.
// On a successful request, the API responds with the modified transfer entity.
// /transfers/:id
// Watch Out!
// You cannot change this setting after the transfer is settled to the linked account.

// =====================================================================================

// Route
// (DM) Route allows you to split payments between third parties, sellers or bank accounts.
// Using Route, you can easily manage settlements, refunds, reconciliations and make vendor payments.
// It is helpful for businesses that disburse payments in a one-to-many model.

// Following is the flow in which the funds move in Route:

// Customer makes a purchase on your site.
// You transfer the funds to the Linked Accounts. You can choose to do any one of the followig:
// Defer the transfer from being settled.
// Define a time until which the transfer settlement should be delayed.
// Razorpay settles funds to the Linked Account and sends a webhook payload to you.

// Use Cases
// The following are the business models across many industries that can benefit by using (DM) Route:
// Financial Services
// A financial service that has multiple insurance aggregators might want to use Route to split payments between the aggregators.
// Acme Corporation is an insurance aggregator.
// A customer logs into the website and purchases a health insurance policy of ₹ 10,000 from Insurer A and an add-on,
// cancer insurance for ₹ 5,000 from Insurer B.
// Acme can split the transaction amount and transfer it to the Linked Accounts of insurers A and B using Route.
// Suppose the customer cancels the add-on policy and seeks a refund. Being a Linked Account,
// Insurer B can directly refund ₹ 5,000 using the Refund Credits from their Dashboard.

// Route
// Delaram Majestic Route enables you to split payments received using the Razorpay Payment Gateway or other products
// (such as Payment Links, Payment Pages, Invoices, Subscriptions and Smart Collect) and transfer the funds to your vendors.

// Multiple Bank Accounts
// A business that has multiple bank accounts might want to route payments to these accounts based on some rules, such as:
// Geographical location.
// Centralize and schedule payouts to bank accounts based on business needs.
// Control settlement schedules for each account.
// For example, Acme Corporation is a business operating in various locations in India,
// such as Chennai, Bangalore, Hyderabad, Lucknow and Delhi. It collects payments centrally
// and disburses payments via bank accounts created for every location. For example,
// suppose the Chennai location makes a sale of ₹ 10,000 for a month,
// and the Hyderabad location makes a sale of ₹ 15,000. In this case, Acme Corporation processes the respective payouts
// and decides the settlement schedules for each location

// Online Marketplace
// An online marketplace is a business platform where products are sold not only by
// you but other third-party sellers as well. In such cases, you, as the marketplace owner,
// need to collect the payment from the customer and disburse it to the third-party sellers.
// For example, Acme Corporation owns an e-commerce site. A customer buys a phone costing
// ₹ 10,000 from seller A, and a phone case costing ₹ 500 from seller B. Acme Corporation collects the total payment (₹ 10,000 + ₹ 500)
// from the customer and then transfers the respective funds to seller A and seller B. Also,
// Acme Corporation collects some charges from these sellers for using the e-commerce platform.

// ================================================================

// The following details are required to create a linked account:

// Primary Details
//   - Linked Account Name
//   - Type
//   - Contact Number

// Bank Details
//   - Account Number
//   - Account Type
//   - IFSC
//   - Beneficiary Name

// Once a linked account is created, it will be assigned an account_id .

// Transfers

// Transfer Requirements
// Below are the requirements to initiate a transfer:
// Your account must have sufficient funds to process the transfer to the linked account.
// The transfer will fail in case of insufficient funds.
// You can only transfer the captured payments.
// You can create more than one transfer on a payment_id. However, the total transfer amount should not exceed the captured payment amount.
// You cannot request a transfer on payment once a refund has been initiated.

// Types of Transfers
// You can transfer funds to linked accounts using one of the following methods:

// @-Transfer via Orders
// You can set up a transfer at the time of order creation.

// 1- you set up transfer using orders API
// 2_ customer make paymetns against the order
// 3_ payment is recieved and transfered to linked accounts
// 4_ linked account receives the transfered amounts

// Watch Out!
// If a Transfer via Order initiated by you fails, we will retry this transfer on consecutive days.
// There will be maximum of 3 retries.

// @-Transfer via Payments
// You can initiate a transfer once the payment has been received from the customer.
// customer makes the payment vis (meli ) gateway or other products
// you recieve the payment and transfer it to linked accounts using meli routes
// linked account recieves the amounts and can view it in their dashboard

// @-Direct Transfer
// You can initiate a transfer directly from existing funds in your (DM) account.

// you initiate a direct transfer from your existing balance , using api
// linked account recieve the amount and can view in their dashboard
//
