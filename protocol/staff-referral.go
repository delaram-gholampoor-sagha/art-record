

package protocol

import "../libgo/protocol"

type StaffReferral interface {
	StaffID() [16]byte    // staff-status domain
	ReferralID() [16]byte // user-status domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
