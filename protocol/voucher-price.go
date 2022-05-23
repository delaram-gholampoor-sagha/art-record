package protocol

import "time"

type VoucherPrice interface {
	VoucherID() [16]byte // voucher domain
	Each() uint8         // Each time use
	Price() int64        // Max discount price per use
	Time() time.Time     // Save time
	RequestID() [16]byte // user-request domain
}

type VoucherPrice_StorageServices interface {
	Save(vp VoucherPrice) error

	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vp VoucherPrice, err error)
	Last(id [16]byte) (vp VoucherPrice, err error)
}
