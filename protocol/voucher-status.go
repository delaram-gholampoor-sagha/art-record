package protocol

import (
	"../libgo/protocol"
)

// Voucher or gift voucher is a bond of the redeemable transaction type which is worth a certain monetary value and
// which may be spent only for specific reasons or on specific products
type VoucherStatus interface {
	VoucherID() [16]byte    //
	Status() Voucher_Status //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type VoucherStatus_StorageServices interface {
	Save(vs VoucherStatus) protocol.Error

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vs VoucherStatus, err protocol.Error)
	Last(voucherID [16]byte) (vs VoucherStatus, numbers uint64, err protocol.Error)
}

type Voucher_Status uint8

const (
	Voucher_Status_Unset Voucher_Status = 0

	Voucher_Status_AutoGenerated Voucher_Status = (1 << iota)

	Voucher_Status_PartlyUsed
	Voucher_Status_FullyUsed

	Voucher_Status_Expired
	Voucher_Status_Blocked
	Voucher_Status_Revoked
)
