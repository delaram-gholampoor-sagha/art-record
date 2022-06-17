package protocol

import (
	"../libgo/protocol"
	"../libgo/time/utc"
)

// ProjectBudget indicate the domain record data fields.
type ProjectAnalytic interface {
	ProjectID() [16]byte // project domain
	Day() utc.DayElapsed

	TotalTasks() uint64      //
	CompletedTasks() uint64  //
	RejectedTasks() uint64   //
	Progress() uint64        //
	TaskOverDueTime() uint64 // due time

	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
