package protocol

import "time"



type VoucherPercentage interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	Percentage() uint64  // Max discount percentage per use
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}


type VoucherPercentage_StorageServices interface {
	Save(vp VoucherPercentage) error
	
	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vp VoucherPercentage, err error)
	Last(id [16]byte) (vp VoucherPercentage, err error)
}