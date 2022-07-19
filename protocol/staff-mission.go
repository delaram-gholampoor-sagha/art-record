/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"
type StaffMission interface {
	StaffID() [16]byte        // staff domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours //
	Type() StaffMission_Type  //
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type StaffMission_Type uint8

const (
	StaffMission_Type_Unset StaffMission_Type = iota
	StaffMission_Type_Meeting
	StaffMission_Type_Negotiation
	StaffMission_Type_Deal
	StaffMission_Type_Study
	// StaffMission_Type_
)
