package protocol

// OrgBudgetEstimationUTC is weekely budget domain
type OrgBudgetEstimationUTC interface {
	Week() utc.WeekElapsed           //
	Income() protocol.AmountOfMoney  //
	Outcome() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}
