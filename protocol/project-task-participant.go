package protocol


// Task indicate the comment record data fields.
type TaskParticipant interface {
	TaskID() [16]byte           // quiddity domain. use to get and show locale title.
	UserID() [16]byte           // user-status domain
	Type() TaskParticipant_Type //
	Time() protocol.Time        // Save time
	RequestID() [16]byte        // user-request domain
}

type TaskParticipant_Type uint8

const (
	TaskParticipant_Type_Unset TaskParticipant_Type = iota
	TaskParticipant_Type_Assigner
	TaskParticipant_Type_Assigned
	TaskParticipant_Type_Reviewer
)
