package protocol

type StaffReferral interface {
	StaffID() [16]byte    // staff domain
	ReferralID() [16]byte // user domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}
