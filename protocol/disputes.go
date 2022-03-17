package protocol

import "time"

type Dispute struct {
	ID uint64 `json:"id"`
	// Indicates the type of entity. In this case, it is dispute.
	Entity string `json:"entity"`
	// Unique identifier of the payment against which the dispute is created.
	Payment_ID uint64 `json:"payment_id"`
	// Amount, in currency subunits, for which the dispute is created.
	Amount uint64 `json:"amount"`
	// 3-letter ISO currency code associated with the amount.
	Currency string `json:"currency"`
	// The amount, in currency subunits, deducted from your Razorpay current balance when the dispute is lost.
	// This amount will be 0 unless the status of dispute is updated to lost.
	Amount_deducted uint64 `json:"amount_deducted"`
	// Unique identifier of the dispute provided by the gateway.
	Gateway_dispute_id string `json:"gateway_dispute_id"`
	// Code associated with the reason for dispute.
	Reason_code string `json:"reason_code"`
	// Reason for raising the dispute.
	Reason_description string `json:"reason_description"`
	// Timestamp (in Unix) by which a response should be sent to the customer.
	Respond_by time.Time `json:"respond_by"`
	// Possible statuses:
	// open: Indicates that the dispute has been created.
	// under_review: Indicates that the dispute is being reviewed by the issuing bank.
	// won: Indicates that the bank has accepted the remedial documents and you have won the chargeback.
	// lost: Indicates that the bank did not accept the remedial documents and you have lost the chargeback.
	// closed: Indicates that the fraudulent transaction is closed after you either provided the details of the transaction or make a refund to the customer.
	// This is seen in fraudulent transactions only.
	Status string `json:"status"`
	// phase
	// string Phase that is associated with the dispute. Possible values:
	// fraud: A dispute raised by the bank when it suspects a transaction to be fraudulent based on the risk analysis.
	// retrieval: A request initiated by the customer with their issuer bank for additional information about a transaction.
	// chargeback: A refund claim initiated by the customers with their issuer banks. In such cases, the bank starts an official inquiry.
	// pre_arbitration: A chargeback that you have won is challenged by the customer for the second time.
	// arbitration: A chargeback that you have won is challenged for a third time by the customer and the card networks directly get involved.
	Phase      string    `json:"Phase"`
	Comments   string    `json:"comments"`
	Created_At time.Time `json:"created_at"`
}

// respond_by
// integer Timestamp (in Unix) by which a response should be sent to the customer.

// About Disputes
// A dispute is a situation where a customer or the issuing bank questions the validity of payments.
// It may arise due to unauthorized charges, failure to deliver the promised merchandise, or excessive charges levied by the business.

// Dispute Phases
// A dispute can be in any of the following phases:

// Phase :  Description

// Fraud : A dispute is raised by the bank when it suspects a transaction to be fraudulent based on the risk analysis.

// Retrieval : A request is initiated by the customer with their issuer bank for additional information about a transaction.
// This is essentially a "soft" chargeback.

// Chargeback : A refund claim initiated by the customers with their issuer banks. In such cases, the bank starts an official inquiry.

// Pre-Arbitration : A chargeback that you have won and is again challenged by the customer for the second time.

// Arbitration : A chargeback that you have won and is rechallenged for the third time by the customer,
// and the card networks directly get involved. These disputes are usually costly.

// Handy Tips : The pre-arbitration and arbitration dispute phases are usually long-drawn, complicated, and challenging.
//It is advised to take remedial action during the retrievals and chargeback phases to avoid complications.

// Dispute States
// A dispute can have any of the following statuses (in the Razorpay system):

// Status :  Description

// Open : Indicates that the dispute has been created.

// Under Review : Indicates that the issuing bank has reviewed the dispute.

// Won : Indicates that the bank has accepted the remedial documents, and you have won the chargeback.

// Lost :  Indicates that the bank did not accept the remedial documents and you have lost the chargeback.

// Closed : Indicates the state when a fraudulent transaction is closed after you provide details of the transaction or make a refund to the customer.
// This is seen in fraudulent transactions only.

// Disputes Process Flow
// Following is the process flow for Disputes:

// A dispute is constantly raised by the issuing bank. However, it can be initiated by the customer as well.

// Initiated by the issuing bank: The issuing bank suspects a fraudulent transaction and asks for your justification.
// .بانک صادرکننده به تراکنش متقلبانه مشکوک شده و از شما درخواست توجیه می کند
// Initiated by the customer: The customer claims that the transaction was unauthorized and raises it with the issuing bank.
// مشتری ادعا می کند که این تراکنش غیرمجاز بوده است و آن را با بانک صادر کننده مطرح می کند.

// =========================================

// You will be notified about the dispute.

//        Accept the dispute. The customer is refunded. In the case of fraud, you must refund the amount.
//        In other cases, your service should auto-refund the amount.
//        Contest the dispute by submitting evidence to prove that the transaction was fair.
//        If you contest, the documents are sent to the customer’s bank. The bank reviews the case and provides a verdict.
//        If you lose the dispute, the amount would be deducted from your account and is sent to the customer.
//        Disputes in International Payments
//        For a dispute raised for international payment, the amount deducted from your account will be based on the day's currency conversion rate when the dispute was created.
//        The currency conversion rate is dependent on the rate charged by processing banks.This conversion rate may vary from that on the day when the payment was created.

// اختلاف را بپذیرید به مشتری بازپرداخت می شود. در صورت کلاهبرداری، باید مبلغ را بازپرداخت کنید.
//          در موارد دیگر، سرویس شما باید مبلغ را به صورت خودکار بازپرداخت کند.
//          با ارائه شواهدی برای اثبات منصفانه بودن معامله، اختلاف را به چالش بکشید.
//          در صورت رقابت، مدارک به بانک مشتری ارسال می شود. بانک پرونده را بررسی می کند و حکم صادر می کند.
//          در صورت از دست دادن اختلاف، مبلغ از حساب شما کسر می شود و برای مشتری ارسال می شود.
//          اختلاف در پرداخت های بین المللی
//          برای اختلافی که برای پرداخت بین‌المللی مطرح می‌شود، مبلغی که از حساب شما کسر می‌شود، براساس نرخ تبدیل ارز در روزی که اختلاف ایجاد شد، خواهد بود.
//          نرخ تبدیل ارز بستگی به نرخی دارد که بانک‌های پردازشگر دریافت می‌کنند. این نرخ تبدیل ممکن است از روزی که پرداخت انجام شده است متفاوت باشد.

// ==========================================


