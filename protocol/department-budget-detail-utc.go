package protocol


type DepartmentBudgetDetailUTC interface {
	DepartmentID() [16]byte            // department domain
	Week() utc.WeekElapsed             //
	Type() DepartmentBudgetDetail_Type //
	Amount() protocol.AmountOfMoney    //
	Time() protocol.Time               // Save time
	RequestID() [16]byte               // user-request domain
}
