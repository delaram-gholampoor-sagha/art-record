package protocol

// ProjectTiming indicate the domain record data fields.
type ProjectTiming interface {
	ProjectID() [16]byte         // project domain
	ExpectedTime() protocol.Time //
	Deadline() protocol.Time     //
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}
