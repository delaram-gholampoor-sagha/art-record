package protocol

type DepartmentBudgetActualUTC interface {
	DepartmentID() [16]byte          // department domain
	Week() utc.WeekElapsed           //
	Income() protocol.AmountOfMoney  //
	Outcome() protocol.AmountOfMoney //
	Time() protocol.Time             // Save time
	RequestID() [16]byte             // user-request domain
}
