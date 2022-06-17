package protocol

// Project indicate the domain record data fields.
type Project interface {
	ProjectID() [16]byte // group domain. use to get and show locale title.

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
