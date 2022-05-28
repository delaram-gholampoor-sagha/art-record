package protocol

import (
	"../libgo/protocol"
	"../libgo/time/earth"
	"../libgo/time/utc"
)

type StaffMission interface {
	StaffID() [16]byte        // staff-status domain
	Type() StaffMission_Type  //
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours //
	Time() protocol.Time      // Save time
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
