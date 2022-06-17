package protocol


// OrgBudgetEstimationUTC is weekly budget domain
type DepartmentBudgetEstimationUTC interface {
	DepartmentID() [16]byte          // department domain
	Week() utc.WeekElapsed           //
	Income() protocol.AmountOfMoney  //
	Outcome() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}
