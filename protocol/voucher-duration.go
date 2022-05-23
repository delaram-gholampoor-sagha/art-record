package protocol

import (
	"time"

	"../libgo/protocol"
)

type VoucherDuration interface {
	VoucherID() [16]byte          // voucher domain
	Each() uint8                  // Each time use
	Epoch() VoucherDuration_Epoch //
	Duration() protocol.Duration  // from Epoch()
	Time() time.Time              // Save time
	RequestID() [16]byte          // user-request domain
}

type VoucherDuration_StorageServices interface {
	Save(vp VoucherPercentage) error

	Count(id [16]byte) (length uint64, err error)
	Get(id [16]byte, versionOffset uint64) (vp VoucherPercentage, err error)
	Last(id [16]byte) (vp VoucherPercentage, err error)
}

type VoucherDuration_Epoch uint8

const (
	VoucherDuration_Epoch_Unset VoucherDuration_Epoch = iota
	VoucherDuration_Epoch_IssueDate
	VoucherDuration_Epoch_LastUse
)
