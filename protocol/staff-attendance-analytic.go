/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// StaffAttendanceAnalytic indicate the domain record data fields.
type StaffAttendanceAnalytic interface {
	StaffID() [16]byte   // staff domain
	Day() utc.DayElapsed //
	Hours() uint64       //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
