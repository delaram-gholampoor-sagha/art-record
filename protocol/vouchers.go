package protocol

import "time"

// vouchers

// primeCost := the direct cost of a commodity in terms of the materials
// and labor involved in its production, excluding fixed costs.

// هزینه مستقیم یک کالا از نظر مواد
//  و نیروی کار درگیر در تولید آن، به استثنای هزینه های ثابت.

type Vouchers_Content struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"user_id"`
	Explantion string `json:"Explanation"`
	// gheymate tamam shode mahsol
	PrimeCost string `json:"prime_cost"`
	// gheymate asli
	PrimePrice string `json:"prime_price"`
	// قیمت معلق
	// we don't know what the price will be yet.
	PendingPrice    string `json:"pending_price"`
	DiscountedPrice string `json:"discounted_price"`
	// enabling or disabling vouchers ?
	State               string    `json:"state"`
	ExpirationDateTime  time.Time `json:"expiration_datetime"`
	Create_Mode         string    `json:"create_mode"`
	Agent               string    `json:"agent"`
	Internal_Reference  string    `json:"internal_reference"`
	Min_Amount          uint64    `json:"min_amount"`
	Taxes_And_Fees      bool      `json:"Taxes_and_fees"`
	Extras              bool      `json:"extras"`
	Value               Value
	Validity_Start_Date time.Time `json:"validity_start_date"`
	Validity_End_Date   time.Time `json:"validity_end_date"`
	Travel_Start_Date   time.Time `json:"travel_start_date"`
	Travel_End_Date     time.Time `json:"travel_end_date"`
	IssueDateTime       time.Time `json:"issue_datetime"`
}

type Value struct {
	ID                uint64 `json:"id"`
	Voucher_Value     uint64 `json:"voucher_value"`
	Remaining_Value   uint64 `json:"remaining_value"`
	Reusable          bool   `json:"reusable"`
	Category_ID       string `json:"category_ID"`
	Product_ID        uint64 `json:"resuable"`
	Apply_Discount_To string `json:"apply_discount_To"`
}

type Voucher interface {
	GetVoucherExplanation(voucherID uint64) (Vouchers_Content, error)
	CreateVoucherByMoney(money string) (Vouchers_Content, error)
	CreateVoucherByPercentage(percentage int) (Vouchers_Content, error)
	GenerateVoucherByTime(expirationTime string, issueDate string) (Vouchers_Content, error)
	GenerateDisposableVoucher(voucher Voucher) (Vouchers_Content, error)
	GetVoucherExpirationDate(voucherID uint64) (Vouchers_Content, error)
	GetVoucherIssueDate(voucherID uint64) (Vouchers_Content, error)
	GetConsumerofVoucher(voucherID uint64) (string, error)
	CheckVoucherValidity(voucherID uint64) (bool, error)
}

// In business-to-business transactions, often the payments are not due immediately.
// They can be paid with an allowed delay that can vary between 30, 60, or 90 days.
// When the company receives the supplies with the invoice, instead of releasing the payment immediately,
// it creates a voucher as a reminder of the payments due or as a statement of the payment already made.

// در تراکنش‌های تجاری، اغلب پرداخت‌ها بلافاصله سررسید نمی‌شوند.
// می توان آنها را با تاخیر مجاز پرداخت کرد که می تواند بین 30، 60 یا 90 روز متفاوت باشد.
// هنگامی که شرکت لوازم را همراه با فاکتور دریافت می کند، به جای اینکه پرداخت را فوراً انجام دهد،
// یک کوپن به عنوان یادآوری پرداخت های سررسید یا به عنوان بیانیه پرداخت قبلاً انجام شده ایجاد می کند.

// A voucher is an internal document within a company
// that is issued by the accounts payable (AP) department.
// It can be seen as a“memorandum” of the liabilities of the company,
// and it is used to authorize a payment.

// The invoice received from the supplier
// فاکتور دریافت شده از تامین کننده
// The data of the supplier to be paid (name, address, telephone number)
// داده های تامین کننده ای که باید پرداخت شود (نام، آدرس، شماره تلفن)
// The data for the payment (amount due, including a possible discount and due date for the payment)
// داده های پرداخت (مبلغ سررسید، از جمله تخفیف احتمالی و تاریخ سررسید برای پرداخت)
// The initial purchase order made by the company
// سفارش خرید اولیه انجام شده توسط شرکت
// The receipt that confirms that the company received the goods or the services stated in the invoice
// رسیدی که تأیید می کند شرکت کالا یا خدمات مندرج در فاکتور را دریافت کرده است
// The general ledger accounts – needed for accounting reasons
// حساب های دفتر کل - به دلایل حسابداری مورد نیاز است
// The signature of an authorized representative at the company (such as the head of the accounts payable department) that validate the // purchase and the payment
// امضای یک نماینده مجاز در شرکت (مانند رئیس بخش حساب های پرداختنی) که خرید و پرداخت را تأیید می کند.
// The proof of payment, which is included in the voucher documentation
// گواهی پرداخت، که در اسناد کوپن موجود است

// ==============================================

// Dashboard
// Market place

// inventory
// products
// extras
// categories
// resources
// vouchers
// promo codes
// pickups

// orders
// customers
// agents
// scheduale
//    @- calender
//    @- manifest
//    @- agenda
// reports
// sell online
// integrations
// settings

// VOUCHERS
// 1_ Manually by the operator
// 2_ Automatically by the guest
// 3_ Automatically by the operator

// Issue new vouchers
// Do you want to generate new codes automatically, or do you want to manually enter or import codes ?
// code list : enter one code per line maximun 20 characters per code.
// CreateMode       =>  Automatic                        // Quantity              // codeslist
//                  =>  Manual
// ==========================================================================================

// set usage and travel Date restrictions

// Validity Date: Optional - Select a From and To date for when this voucher will be valid for use.
// Leaving the To blank will create an open ended voucher, valid from the From date.
// maybe i want to create these vouchers only on valentines day so i can set the date here
// validity date :		    from :                        to :

// allows you to set restrictions on when the vouchers can be redeemed for
// Travel Date: Optional - Select a From and To date for when the customer needs to travel between.
// Leaving this blank will create a non restrictive voucher.
// travel date :             from :                        to :

// Can only be redeemed on: Optional - Restricts the redemption to certain days of the week.
// can only be redeemed on :        MON    TUE     WED    THU     FRI    SAT     SUN

// Agent: This field is to help you track agents who are using your vouchers.
// Select your agent if you are generating a voucher for a specific agent and want this agent to appear on orders where
// this voucher is redeemed by your customers via your website.
// Agent:        e .g : Digi kala

// Internal Reference: Add a reference to easily report, search and delete vouchers.
// This will not be visible to customers. If left blank, the system will auto generate a number.
// this helps you delete or export certain group of vouchers later on
// if we leave this blank it will use a timestamp be default to help you search it by time.
// Internal Reference :         => find me later

// Min Amount: Enter a minimum booking amount.
// Upon checkout, the Customers shopping cart must equal or greater than this amount before they can use the voucher code.
// your users shopping cart must be above this amount for the voucher code to be valid
// Min Amount :

// Include all taxes & fees: If checked, taxes will be calculated on the discounted total.
// If not checked, taxes will be calculated on the full amount without discount.
// taxes and fees by default has been ticked which means taxed and fees will be calculated on the remaining amount after discount
// Taxes and fees :           checkbox

// Extras:     checkbox

// here you select a product and add the quantities or maybe select a category and say which category this voucher is valid for ?
//
//
// if you want this voucher to be used as a payment method for a fixed amount and a specific product.
// value :            => fixed amount for one product
//                           @- voucher value         enter the value
//                           @- reusable              check box
//                           @- product               select
//                                                      - intoxicating  thai (thai-class)
//                           apply discount to        @ apply to everyone
//                                                    @ Only to :
//                                                          @ adult
//                                                          @ child
// if you want this voucher to be used as a payment method for a fixed value for the whole shopping cart.
//                    => fixed amount for any product
//                           @- voucher value         enter the value
//                           @- reusable              check box
// this will show up on the check out page as you are processing your order
//                           @- remaining value              select
// if you want this voucher to be used as a payment method for a fixed amount and limit its usage to a specific catalog
//                    => fixed amount for any product within a category
//                           @- voucher value         enter the value
//                           @- reusable              check box
//                           @- category              select
// if you want this voucher to be used as a payment method for a specific product only.
//                    => free product
//                           @- product             select a product
//

// Internal Notes: Add any notes for internal purposes.
// Internal notes :

// voucher code          internal refrence              select agent          all vouchers(based on status)         	SEARCH
// -----------------------------------------------------------------------------------------------------------------------------

//   CODE      STATUS          VALUE                     VALIDITY_DATE        TRAVEL_DATE        AGENT          INTERNAL_REFRENCE
//   VYD22    redeemed    free_product(korean food)      from 17/2/22                           snap food       21654305132065113
//   VPCA5     issued     free_product(dress)        16/10/20  to 16/10/21                      digikala        16461316545313516
// 	 DJLNI    expired     free_product(korean food)  16/10/20  to 16/1021                      sofi restuarant  32164645454545545
//   QPMOE     issued      $200.00 remaining:$200.00     from 16/10/20                         meli bank        32165546465464655

// your guests can automatically create vouchers by purchasing one your products as gift or if we have set up gift cart products
// when a product is purchased as gift or a gift product has been board a voucher will automatically be generated for the value amount that was just purchased.
// so taking a look at an order for a product the travel date section has been replaced with genereratd vouchers

// generated vouchers as payment method ????!!!!!!

// -How-to-Manually-Redeem-a-Voucher
// How to redeem a voucher on behalf of a customer
// From the Orders page, click into the customer order (or create a new internal order):

// Scroll down to the Customer Payment section
// Select Voucher from the Payment Type
// Start typing the Voucher number, and choose from the list of automatically found vouchers.
// Click Redeem

// This will automatically enter the value that the Voucher is valid for.

// I had a power outage and i wasn't able to run my cooking classes but instead of cancelling all my orders i can simply let
// my guests know they are about ot recieve a voucher code valid for the full value and they can jump back online and choose another date using their voucher code to complete their order
// from the agenda first select the day and it will automaticllay select all the orders then click the update button (update order)
// then this dialogue will show up
// UPDATE ORDERS
// what updates would like to make ?
// CHANGE ORDER STATUS TO :
// cancelled
// confirmed
// pending customer
// pending supplier
// on hold

// refund customers
// refund all card payments
// create and apply vouchers

// email customers:
// do not send email ?
// this will automaticlly send the gift cart email to the main customers contact within the order

// Additional Filters - More Options
// The additional filters allows you to filter in specific booking criteria,
// for example orders created online only, with a specific order status or booked by a specific agent.

// All promo codes: filter in a specific promo code that was used on an order
// All my agents: filter in a specific reseller or agent
// All users: filter in orders created by a specific user of your Rezdy account
// All categories: filter in a specific category of products
// All products: filter in a specific product
// All product types: filter in a specific type of product e.g. Day Tour or Rental
// Payment status: filter in a specific payment status e.g. Paid or Refunded.
// All Orders: filter in unreconciled or reconciled orders
// All Orders: filter in direct or distributed orders.
// Booking system: filter in a specific booking system

// As the supplier vs As the agent
// you should allows operators to act as an agent,
// reselling other operator tours and activities within the Marketplace.
//If you have sold other supplier products, you can click the As the agent button and run the same Sales - Orders report as per the above steps.
