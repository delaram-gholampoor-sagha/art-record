/* For license and copyright information please see LEGAL file in repository */

package protocol

import "time"

// VoucherProduct restrict use voucher on specific product or belong to specific category.
type VoucherProduct interface {
	VoucherID() [16]byte  // voucher domain
	Each() uint8          // Each time use
	QuiddityID() [16]byte // quiddity domain. specific product or specific category.
	Time() time.Time      // Save time
	RequestID() [16]byte  // user-request domain
}


type VoucherProduct_StorageServices interface {
	Save(vu VoucherUser) error
	
	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vu VoucherUser, err error)
	Last(id [16]byte) (vu VoucherUser, err error)

	FindVoucherByQuiddityID(quiddityID [16]byte) (vu VoucherProduct , err error)
}
