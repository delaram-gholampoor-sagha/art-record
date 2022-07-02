package protocol

// VoucherUser indicate the domain record data fields.
// Use to restrict the voucher usage by specific users.
type VoucherUser interface {
	VoucherID() [16]byte // voucher domain
	UserID() [16]byte    // user domain.
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type VoucherUser_StorageServices interface {
	Save(vu VoucherUser) (numbers uint64, err protocol.Error)

	Count(voucherID [16]byte) (numbers uint64, err protocol.Error)
	Get(voucherID [16]byte, versionOffset uint64) (vu VoucherUser, numbers uint64, err protocol.Error)

	FindByUser(userID [16]byte, offset, limit uint64) (voucherIDs [][16]byte, numbers uint64, err protocol.Error)
}
