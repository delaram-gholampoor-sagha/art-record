package protocol

// ProjectTaskTag indicate the domain record data fields.
type ProjectTaskTag interface {
	TaskID() [16]byte    // task domain
	Tag() [16]byte       // project-tags domain
	Time() protocol.Time // Save time
}
