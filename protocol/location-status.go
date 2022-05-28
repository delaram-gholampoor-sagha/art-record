package protocol

import "../libgo/protocol"

type LocationStatus interface {
	LocationID() [16]byte    // location domain
	Status() Location_Status //
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain

}

type Location_Status uint8

const (
	Location_Status_Unset Location_Status = iota
	Location_Status_Full
	Location_Status_Empty
	Location_Status_UnUsable
	Location_Status_Busy
	Location_Status_
)
