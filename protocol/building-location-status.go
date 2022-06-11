package protocol

type BuildingLocationStatus interface {
	BuildingLocationID() [16]byte // quiddity domain
	Status() Location_Status      //
	Time() protocol.Time          // Save time
	RequestID() [16]byte          // user-request domain
}

type Location_Status uint8

const (
	Location_Status_Unset           Location_Status = 0
	Location_Status_PermanentClosed Location_Status = (1 << iota)
	Location_Status_TemporaryClosed
	Location_Status_Blocked
)
