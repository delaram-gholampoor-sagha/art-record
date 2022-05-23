package protocol

import "time"



type VoucherInvoice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	MinPrice() int64     // Minimum invoice price
	MinAmount() uint64   // Minimum invoice product numbers
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}




type VoucherInvoice_StorageServices interface {
	Save(vp VoucherPercentage) error
	
	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vp VoucherPercentage, err error)
	Last(id [16]byte) (vp VoucherPercentage, err error)
}