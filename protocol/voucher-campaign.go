package protocol

// VoucherUser restrict voucher usage by specific user or user type.
type VoucherCampaign interface {
	VoucherID() [16]byte  // voucher domain
	CampaignID() [16]byte // marketing-campaign domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}
