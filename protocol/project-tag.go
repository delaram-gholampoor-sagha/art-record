package protocol

// ProjectList indicate the domain record data fields.
type ProjectTag interface {
	ProjectID() [16]byte // Project domain
	Tag() [16]byte       // project-tags domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
