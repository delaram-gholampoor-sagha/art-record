package protocol

// VoucherUser restrict voucher usage by specific user or user type.
type VoucherUser interface {
	VoucherID() [16]byte         // voucher domain
	UserID() [16]byte            // user domain.
	UserType() protocol.UserType //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}
