package protocol

import (
	"time"
)

type Invoice_Content struct {
	ID uint64 `json:"id"`
	// or userID ? I want to indicate to the buyer of this product
	UserID string `json:"acossiated_user"`
	// events invoice
	Seller string `json:"seller"`
	// events invoice
	Card_Number string `json:"card_number"`
	Price       string `json:"price"`
	Tax         string `json:"tax"`
	SumـTotal   string `json:"sum_total"`
	// سود شما از این خرید
	Profit_From_Purchase string `json:"profit_from_purchase"`
	// تعداد کالا
	Number_of_Goods string `json:"number_of_goods"`
	// قیمت اعمال شده
	Price_Applied string `json:"price_applied"`
	// or voucherID ?
	VoucherID uint64 `json:"voucher"`
	// The unique identifier of the order associated with the invoice.
	Order_ID uint64 `json:"order_id"`
	// Invoice States
	// An invoice starts in the draft state and moves through several different states in its life cycle.

	// Invoice States :
	// The table below lists the various states of an invoice and gives a brief description of each state:

	// draft : Indicates that the invoice has been created and saved. An invoice in this state is not issued and can be edited later.
	// Customer and item fields are required to save an invoice. A saved invoice can be deleted.
	// You cannot view a deleted draft invoice.

	// issued : Indicates that an invoice has been finalized and issued to the customer. An issued invoice cannot be deleted, but can be cancelled.
	// You can edit the following information even after sending it to the customer:
	// - Invoice #
	// - Expiry Date
	// - Customer Notes
	// - Terms and Conditions
	// - Enable or disable partial payment. Know more about issuing an Invoice.

	// partially paid : Indicates that the customer has made a partial payment against an invoice.
	// After a payment has been made, you can neither delete nor cancel the invoice.
	// You can only view the invoice and add internal notes to it.

	// paid : Indicates that the customer has fully paid the invoice amount.
	// Paid is the final state of an invoice. A paid invoice can only be viewed.
	// You can neither delete nor cancel it.

	// expired : Indicates that an invoice has expired. Once expired,
	// the customer cannot make any payments against the invoice,
	// and you cannot cancel or delete the invoice. You can view an expired invoice.

	// cancelled : Indicates that you have cancelled the invoice. Though you can view the cancelled invoice,
	// the customer cannot view the invoice or make payments against the invoice.
	// You can only cancel an unpaid issued invoice.

	// deleted : This state indicates that the invoice has been deleted.
	// You can only delete an invoice in the draft state and not in any other state.
	State      string `json:"state"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	// ================================
}

type Invoicee struct {
	ID uint64 `json:"id"`
	// Here, it is invoice.
	Entity string `json:"entity"`
	// Unique number you added for internal reference
	Invoice_Number uint64 `json:"invoice_number"`
	Customer_ID    uint64 `json:"customer_id"`
	// Details of the line item that is billed in the invoice. Maximum of 50 line items.
	Line_Item string `json:"line_item"`
	// Unique identifier of a payment made against this invoice.
	Payment_ID string `json:"payment_id"`
	// Timestamp, in Unix format, at which the invoice will expire.
	Expire_by time.Time `json:"expire_by"`
	// Timestamp, in Unix format, at which the invoice was issued to the customer.
	Issue_at time.Time `json:"issue_at"`
	// Timestamp, in Unix format, at which the payment was made.
	Paid_At time.Time `json:"paid_at"`
	// Timestamp, in Unix format, at which the invoice was cancelled.
	Cancelled_At time.Time `json:"cancelled_at"`
	// Timestamp, in Unix format, at which the invoice expired.
	Expired_At time.Time `json:"expire_at"`
	// The delivery status of the SMS notification for the invoice sent to the customer.
	// Possible values:
	// pending
	// sent
	SMS_Status string `json:"sms_status"`
	// 	The delivery status of the email notification for the invoice sent to the customer.
	// Possible values:
	// pending
	// sent
	Email_Status string `json:"email_status"`
	// boolean Indicates whether the customer can make a partial payment on the invoice. Possible values:
	// true: The customer can make partial payments.
	// false (default): The customer cannot make partial payments.
	Partial_Payment bool `json:"partial_payment"`
	// Amount to be paid using the invoice. Must be in the smallest unit of the currency
	Amount uint64 `json:"amount"`
	// Amount paid by the customer against the invoice.
	Amount_Paid uint64 `json:"amount_paid"`
	// The currency associated with the invoice. You must mandatorily pass this parameter if accepting international payments.
	// If you have passed currency as a sub-parameter in the line_item object, you must ensure that the same currency is passed in both places.
	Currency string `json:"currency"`
	// A brief description of the invoice.
	Description string `json:"description"`
	// Any custom notes added to the invoice. Maximum of 2048 characters.
	Notes string `json:"notes"`
	// The short URL that is generated. This is the link that can be shared with the customer to receive payments.
	Short_URL string `json:"short_url"`
	// Here, it is invoice.
	Type string `json:"type"`
	// Timestamp, in Unix format, that indicates the issue date of the invoice.
	Date string `json:"date"`
	// Any terms to be included in the invoice. Maximum of 2048 characters.
	Terms string `json:"terms"`
	// Any comments to be added in the invoice. Maximum of 2048 characters.
	Comment string `json:"comment"`
}



type Line_Item struct {
	ID uint64 `json:"id"`
	// Unique identifier of the item generated using Items API that has been billed in the invoice.
	Item_ID     uint64 `json:"item_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// The price of the item.
	Amount uint64 `json:"amount"`
	// The currency associated with the item.
	Currency string `json:"currency"`
	// Here, it is invoice.
	Type string `json:"type"`
	// The quantity of the item billed in the invoice. Defaults to 1.
	Quantity string `json:"quantity"`
}

// shipping_address
// Details of the customer's shipping address.
type Shipping_Address struct {
	ID uint64 `json:"id"`
	// The customer address type. Here it is shipping_address.
	Type string `json:"type"`
	// Defines if this is the primary address.
	Primary bool `json:"primary"`
	// The first line of the customer's address.
	Line1 string `json:"line1"`
	// The second line of the customer's address.
	Line2   string `json:"line2"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	State   string `json:"state"`
	Country string `json:"country"`
}

// billing_address
// Details of the customer's billing address.
type Billing_Address struct {
	ID uint64 `json:"id"`
	// The customer address type. Here it is shipping_address.
	Type string `json:"type"`
	// Defines if this is the primary address.
	Primary bool `json:"primary"`
	// The first line of the customer's address.
	Line1 string `json:"line1"`
	// The second line of the customer's address.
	Line2   string `json:"line2"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	State   string `json:"state"`
	Country string `json:"country"`
}


// we'll talk about it later ... maybe another microservice
// امتیاز کسب شده از این خرید
// Points earned from this purchase ??

type Invoice interface {
	// Invoice ??
	GetInvoice(invoiceID uint64) (Invoice_Content, error)
	GetInvoices() ([]Invoice_Content, error)
	GetInvoicesByAdmin(userId uint64, invoiceID uint64) (Invoice_Content, error)
	RegisterInvoice(invoice Invoice_Content) (Invoice_Content, error)

	// Find Invoice
	FindInvoiceByProductID(productID uint64) (Invoice_Content, error)
	FindInvoiceByUserID(userID uint64) (Invoice_Content, error)
	FindInvoiceByAcossiatedProductOwner(productID uint64) (Invoice_Content, error)
	FindInvoiceByDateTime(time string) (Invoice_Content, error)
	FindInvoiceByState(state string) (Invoice_Content, error)
}

// what-is-invoice-processing
// Here are the invoice processing steps:
// Capture. Vendor invoices are received by fax, mail, email, or captured by an accounts payable system.
// Register. Invoices are prepared for approval by coding and updating the invoice data.
// Dispatch. Authorized approvers are sent invoices for a payment decision.
// Approval/Rejection. Approvers will either approve or reject the dispatched invoice for payment.
// Payment. The approved invoices are then processed for payment.
// Archive. Paid invoices are stored either in a file cabinet or online system with invoice processing history for later reference or audits.

// https://www.stampli.com/invoice-management/
// invoice status : invoice details have been registered
// handled by : waiting to be assigned
// processing began : 08/21/2018 recieved from bethony glover
// due date : 09/20/2018 change ?!
// financial system : export
// invoice type : invoice change ?!

// invoice NO : 1654613513551
// vendor : Corpo industries INC
// invoice date : 08/07/2018

// total to be paid : 1951.00

// company code : Hone comp holding Inc
// purchase order : POU87686856

// GL accounts : add/edit line items
//     GL sccount : 130 millescenoues expense
//     class :    isccelleniuos
//     memo :
//     location
// description

// https://coreintegrator.com/invoice-processing-steps/
// These are the general steps to invoice processing:

// Receive Invoice
// Copy and File Invoice
// Invoice Approval
// Invoice Payment

// Automated Invoice Approval
// Once the accounting department reviews the invoice data, the system routes invoices for approval based on your business’ specific invoice approval rules.

// And you can easily conﬁgure the approval matrix based on a multitude of rules, such as:
// Person
// Amount
// Location
// Invoice type
// Vendor

// https://www.freshbooks.com/hub/invoicing/processing-invoices
// How to Process an Invoice: A Guide for Small Business Owners
// Step 1: Verifying and Tracking Information
// A purchasing company needs to verify the purchase, ensure correct payment and deliver the payment within the agreed upon terms.
// Invoices should include the following information to help the vendor and purchaser track their expenses or inventory and update their financial records:
// Date the vendor created and sent the invoice.
// Contact information of both the vendor and the purchaser, particularly billing information and point of contact.
// Purchase details, including product or service details and pricing.
// Payment information

// Step 2: Data Entry and General Ledger Coding
// Once the AP staff verifies that the vendor invoice contains all the correct information,
// they need to enter the data manually or using an automation tool and code it for accounting purposes.
// General Ledger Coding refers to a coding system that makes it easier to track debits and credits.

// Manually entering this data can take a lot of staff time and carries the risk of human error,
// which can be detrimental to a company’s financial records. Using an automated system can reduce
// invoice-processing costs by 75% to 85% while decreasing errors at the same time.
// Having the tools you need to track this data also improves access to invoice data,
// which improves the service to vendors and results in the faster turnaround on payments

// Step 3: Forwarding and Receiving Approval
// After an AP department verifies invoice information,
// it needs to submit the invoice for approval before they can send a payment.
// A slow approval process can have a significant payment turnaround times and revenue.
// Paper invoices can sit on a busy employee’s desk or get misplaced as it is moved around from desk to desk.
// By using an automated invoice processing system, an AP department can save time tracking down lost documents or requesting invoice copies from the vendor.

// These solutions digitally capture the data from paper and electronic invoices and put them through
// a custom-designed workflow that speeds up the entire approval process.
// Reducing or eliminating the need for paper invoices will lower outgoing costs.
// According to experts, the cost of a paper invoice can range between $12 to $30 to process with an average cost close to $15.

// While larger companies with a more complex accounts payable process can cost
// nearly $40 per invoice. Online automated invoicing cost significantly less at about
// $3.50 per invoice process. Automation can save your company hundreds of thousands of dollars per year.

// How to enter an invoice into Accounts Payable?
// After an invoice is reviewed and approved it can be entered into the Accounts Payable.
// After a vendor invoice has been approved, the recording of the invoice will include:

// a credit to Accounts Payable
// a minimum of one debit to another account. The debit amount usually involves one of the following:
// an expense (Repairs & Maintenance Expense, Advertising Expense, Rent Expense, etc.)
// a prepaid asset (Prepaid Expenses, Prepaid Insurance)
// a fixed or plant asset (Equipment, Fixtures, Vehicles, etc.)

// https://www.procuredesk.com/faster-invoice-processing/
// What is Invoice processing?
// Invoice processing (referred to as the invoice reconciliation process) also is defined as a process followed by
// the Accounts payable team to upload the invoice, match the invoice with a purchase order, or a lack thereof.
// The steps could vary based on your company process, but in general, here are the steps for invoice processing.

// Key in the invoice in the system.
// Match the invoice with the purchase order, assuming you have a purchase order process in place.
// Resolve the exceptions in the matching process.
// In case there is no purchase order, find out who ordered the product or service from the vendor.
// Then route the invoice for approval to the appropriate person. In case they haven’t approved in time, follow up multiple times till they approve!

// Invoice exception review process
// The unit price on the invoice does not match what is on the purchase order.
// The quantity on the invoice does not match what is on the purchase order.
// The invoice doesn’t have a purchase order and it needs to be approved before it can be paid.
// Whatever the reason for the exception, when there is an exception, the Accounts payable team needs to spend time on resolving those exceptions and that takes time.

// =================================================================

// Invoices
// An Invoice is a digital document that summarises the details of an order or a transaction
// and allows customers to initiate payments. A typical invoice contains sale transaction information
// such as the name of the ordered products or services, quantities, price breakup, receipt number, customer information and so on.



// Issue an Invoice
// Only an invoice in the draft state can be issued. 
// Its response is the invoice entity. Its status now would be issued and it will have a short_url generated.
// Also, SMS and email would be sent to the customer based on what parameters were sent initially during creation.


// Send Notifications
// You can send notifications with the short URL to the customer via email or SMS using the following endpoint:

// /invoices/:id/notify_by/:medium
// Path Parameters
// id : The unique identifier of the invoice whose link is to be sent by SMS or email.
// mandatory

// medium : string Possible values:
// sms
// email
// mandatory


