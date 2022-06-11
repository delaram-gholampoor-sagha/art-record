package protocol


type PictureStatus interface {
	ObjectID() [16]byte     //
	Status() Picture_Status //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}

type Picture_Status uint8

const (
	Picture_Status_Unset Picture_Status = iota
	Picture_Status_Registered
	Picture_Status_Hidden
	Picture_Status_Deleted
	Picture_Status_Blocked
)
