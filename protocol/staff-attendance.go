/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// StaffAttendance indicate the domain record data fields.
type StaffAttendance interface {
	StaffID() [16]byte          // staff domain
	Type() StaffAttendance_Type //
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type StaffAttendance_Type uint8

const (
	StaffAttendance_Type_Unset StaffAttendance_Type = iota
	StaffAttendance_Type_Enter
	StaffAttendance_Type_Exit
	StaffAttendance_Type_ManualEnter
	StaffAttendance_Type_ManualExit
	StaffAttendance_Type_Start
	StaffAttendance_Type_End
)
