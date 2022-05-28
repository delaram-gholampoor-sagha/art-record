package protocol

// Communicate is the structure of private message to person or groups.
type Communicate struct {
	MessageID [16]byte
	From      [16]byte // UserID
	To        [16]byte // UserID || GroupID
}

type Message interface {
	ID() [16]byte
	Time() Time

	// طبقه بندی - عادی/محرمانه
	// فوریت

	RequestID() [16]byte // user-request domain
}
