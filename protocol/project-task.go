package protocol


// ProjectTask indicate the domain record data fields.
type ProjectTask interface {
	TaskID() [16]byte    //
	ProjectID() [16]byte // project domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
