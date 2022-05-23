/* For license and copyright information please see LEGAL file in repository */

package protocol

import "time"

// VoucherUser restrict voucher usage by specific user or user type.
type VoucherUser interface {
	VoucherID() [16]byte         // voucher domain
	UserID() [16]byte            // user-status domain.
	UserType() protocol.UserType //
	Time() time.Time             // Save time
	RequestID() [16]byte         // user-request domain
}



type VoucherUser_StorageServices interface {
    Save(vu VoucherUser) error
	
	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vu VoucherUser, err error)
	Last(id [16]byte) (vu VoucherUser, err error)

	ListVoucherByUserID(userID [16]byte) (voucherID [16]byte , err error)
	FindUserByVoucherID(voucherID [16]byte) (userID [16]byte , err error)
}

