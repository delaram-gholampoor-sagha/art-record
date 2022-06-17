package protocol

type VoucherDuration interface {
	VoucherID() [16]byte          // voucher domain
	Each() uint8                  // Each time use
	Epoch() VoucherDuration_Epoch //
	Duration() protocol.Duration  // from Epoch()
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type VoucherDuration_Epoch uint8

const (
	VoucherDuration_Epoch_Unset VoucherDuration_Epoch = iota
	VoucherDuration_Epoch_IssueDate
	VoucherDuration_Epoch_LastUse
)
