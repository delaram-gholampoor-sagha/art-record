package protocol

// ProjectTaskTiming indicate the domain record data fields.
type ProjectTaskTiming interface {
	TaskID() [16]byte            // project-task domain
	ExpectedTime() protocol.Time //
	Deadline() protocol.Time     //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}
