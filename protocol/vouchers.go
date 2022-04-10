package protocol

import "time"

type Voucher struct {
	ID uint64
	// who can use it ?
	// UserID    uint64
	UserType  uint8
	RelatedID [16]byte
	// QuiddityID  [16]byte
	// Category_ID string
	// Product_ID  uint64
	Status       uint8
	Create_Mode string
	Min_Amount  uint64

	// Question : can we let the user use voucher as a ticker where its value can be changed?
	// or we can bring this concept to another microservice ?
	// like what we talked about the concept of auction and vouchers.
	Voucher_Value uint64
	// Remaining_Value uint64
	
	Reusable      bool
	// level      string
	// Apply_Discount_To   string
	Validity_Start_Date time.Time
	Validity_End_Date   time.Time
	IssueDateTime       time.Time
}

type Voucherservices interface {
	GetVoucherExplanation(voucherID uint64) (Voucher, error)
	// CreateVoucherByMoney(money string) (Voucher, error)
	// CreateVoucherByPercentage(percentage int) (Voucher, error)
	// GenerateVoucherByTime(expirationTime string, issueDate string) (Voucher, error)
	// GenerateDisposableVoucher(voucher Voucher) (Voucher, error)
	GetVoucherExpirationDate(voucherID uint64) (Voucher, error)
	GetVoucherIssueDate(voucherID uint64) (Voucher, error)
	GetConsumerofVoucher(voucherID uint64) (string, error)
	CheckVoucherValidity(voucherID uint64) (bool, error)
}
