package protocol

// ProjectBudget indicate the domain record data fields.
type ProjectBudget interface {
	ProjectID() [16]byte             // department domain
	Week() utc.WeekElapsed           //
	Expense() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}
