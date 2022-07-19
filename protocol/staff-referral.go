/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type StaffReferral interface {
	StaffID() [16]byte    // staff domain
	ReferralID() [16]byte // user domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}
