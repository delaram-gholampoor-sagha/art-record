package protocol

// RoleSalary is fixed negotiated salary that don't follow org||department||role salary rules.
type RoleSalary interface {
	RoleID() [16]byte                     // role domain
	WeeklySalary() protocol.AmountOfMoney //
	Time() protocol.Time                  // save time
	RequestID() [16]byte                  // user-request domain
}
