package protocol

import "../libgo/protocol"

type StaffStatus interface {
	StaffID() [16]byte    // user-status domain
	Status() Staff_Status //
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type Staff_Status uint8

const (
	Staff_Status_Unset Staff_Status = iota
	Staff_Status_Hire
	Staff_Status_Fire
	Staff_Status_End      // death, end of contract, ...
	Staff_Status_Active   // Shift
	Staff_Status_Inactive // Shift
)
