package protocol

import "time"

type Refunds struct {
	ID     uint64 `json:"id"`
	Entity string `json:"entity"`
	// The amount to be refunded (in the smallest unit of currency).
	Amount uint64 `json:"amount"`
	// The currency of payment amount for which refund is initiated. Here is the list of supported currencies.
	Currency string `json:"currency"`
	// Unique identifier of the payment for which the refund is initiated.
	Payment_ID uint64 `json:"payment_id"`
	// Speed with which the refund is to be processed.
	// Possible values:
	// normal: Indicates that the refund will be processed via the normal speed. That is, the refund will take 5-7 working days.
	// optimum: Indicates that the refund will be processed at an optimal speed based on Razorpay's internal fund transfer logic. That is:
	// If the refund can be processed instantly, Razorpay will do so, irrespective of the payment method used to make the payment.
	// If an instant refund is not possible, Razorpay will initiate a refund that is processed at the normal speed.
	// For example, in case of payments made using debit cards, netbanking or unsupported credit cards.
	Speed      string    `json:"speed"`
	Created_At time.Time `json:"created_at"`
	// This parameter is populated if the refund was created as part of a batch upload. For example, batch_00000000000001
	Batch_ID uint64 `json:"batch_id"`
	// json object Key-value store for storing your reference data. A maximum of 15 key-value pairs can be included.
	Notes string `json:"notes"`
	// A unique identifier provided by you for your internal reference.
	Receipt string `json:"receipt"`
	//  A dynamic array consisting of a unique reference number (either RRN, ARN or UTR) that is provided by the banking partner when a refund is processed.
	// This reference number can be used by the customer to track the status of the refund with the bank.
	// the type of this attribute is an array
	Acquirer_Data string `json:"acquired_data"`
	// Possible values:
	// pending: This state indicates that Razorpay is attempting to process the refund.
	// processed: This is the final state of the refund.
	// failed: A refund can attain the failed state in the following scenarios:
	// Normal refund not possible for a payment which is more than 6 months old.
	// Instant Refund can sometimes fail because of customer's account or bank-related issues.
	Status string `json:"status"`
	// The processing mode of the refund seen in the refund response.
	// // This attribute is seen in the refund response only if the speed parameter is set in the refund request.
	// Possible values:
	// normal: Indicates that the refund will be processed via the normal speed. That is, the refund will take 5-7 working days.
	// optimum: Indicates that the refund will be processed at an optimal speed based on Razorpay's internal fund transfer logic. That is:
	// If the refund can be processed instantly, Razorpay will do so, irrespective of the payment method used to make the payment.
	// If an instant refund is not possible, Razorpay will initiate a refund that is processed at the normal speed.
	Speed_Requested string `json:"speed_requested"`
	// This is a parameter in the response which describes the mode used to process a refund.
	// This attribute is seen in the refund response only if the speed parameter is set in the refund request.
	// Possible values:
	// instant: This means that the refund has been processed instantly via fund transfer.
	// normal: This means that the refund has been processed by the payment processing partner. That is, the refund will take 5-7 working days.
	Speed_Processed string `json:"speed_requested"`
}



// About Refunds
//  There could be situations when customers request a refund of the payments made for 
//  the products or services purchased or availed on your website or app.


// Refunds
//  You can make full or partial refunds to customers. While issuing refunds, 
//  you can choose to process the refunds instantly or at normal speed (within 5-7 working days).


// Refunds Can be Made Only on Captured Payments
//  You can initiate refunds only on those payments that are in captured state. 
//  A payment in authorized state is auto-refunded if not captured within 5 days of creation.


// Refund Types
//   Following are the various types of refunds that you can use to refund payments to your customers:

//     Normal Refund
//     Amount is refunded within 5-7 working days.

//     Instant Refund
//     Amount is refunded almost immediately.
//     By issuing instant refunds to your customers,
//     you can provide a better user experience for them. This also helps in improving their reliability and trust in your business.

//     Batch Refund
//     Issue refunds in bulk using an XLSX or CSV file. 
//     Once you upload a file, it is picked up for processing after 70 minutes. 
//     You can cancel a batch upload in the 70 minutes before it is picked up for processing.


// Handle Refund Chargeback 
//  For the prevention of chargebacks, Razorpay only does source refunds. 
//  It means that money is refunded to the payment method that the customer used to make the payment.
//  For example, if a credit card was used to make the payment, the refund is pushed to the same credit card. 
//  Similarly, in the case of UPI payments, the refund is pushed to the VPA used while making the payment.

//  If a chargeback is received for an instantly refunded payment,
//  the processed refund will have a UTR (Unique Transfer Reference) in the callback. 
//  This UTR appears against the ARN (Application Reference Number) parameter in the Refund entity.
//  The UTR serves as a proof of refund completed between you and Delaram-Majestic Service.

//  Additionally, (Delaram Majestic service) passes the DRN (Delaram Reference Number) of the payment in the Fund Transfer Request sent for the refund.
//  This ties the instant refund back to the parent payment, thereby, serving as a proof of the refund.
//  This data can also be used as a defense against a future chargeback or arbitration case.

// =========================================================================================
// Refund Communication
// If a customer is asking for a refund, the banking partner provides a unique reference number (either RRN, ARN or UTR) when a refund is processed.
// The customer can use this reference number to track the refund status with the bank.
// As a customer, you will be notified of the specific payment to be refunded in 7-10 working days.

// 1. What should I do if I have not received my refund even after 10-12 working days?
// Ideally, refunds are credited in 10-12 working days. In case you have not received your refund,

// 2. How do I check my order status after the amount is deducted from my account?
// Upon a successful transaction, you will receive a confirmation email from Razorpay.

// 3. How do I cancel my order?
// All the queries related to full or partial cancellation of orders should be routed to the business.

// 4. Can I file a complaint if I have not received my order?
// We suggest that you wait for 7-10 business days for physical goods and 1 business day for digital goods.
// If you have still not received the items, or if the items received are materially different, defective, or damaged, contact the business to resolve the issue.
// Include the following details while sending an email to the business:

//    Date of transaction.
//    Amount of transaction.
//    Order ID shared by the business.
//    Payment ID shared by Delaram-Majestic(pay_9uxxxxxxx34z).
//    Description of the problem.

// ======================================================
// Fetch Multiple Refunds for a Payment
// The following endpoint retrieves multiple refunds for a payment. By default, only the last 10 refunds are returned.
// You can use count and skip parameters to change that behavior.
// /payments/:id/refunds

// Fetch a Specific Refund for a Payment
// The following endpoint retrieves details of a specific refund made for a payment.
// /payments/:payment_id/refunds/:refund_id










