package protocol



type AreaStatus interface {
	AreaID() [16]byte    // quiddity domain
	Status() Area_Status //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

type Area_Status uint8

const (
	Area_Status_Unset           Area_Status = 0
	Area_Status_PermanentClosed Area_Status = (1 << iota)
	Area_Status_TemporaryClosed
	Area_Status_Blocked
)
