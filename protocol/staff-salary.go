package protocol

import (
	"../libgo/protocol"
)

// StaffSalary is fixed negotiated salary that don't follow org||department||role salary rules.
type StaffSalary interface {
	StaffID() [16]byte                    //
	WeeklySalary() protocol.AmountOfMoney //
	Time() protocol.Time                  // Save time
	RequestID() [16]byte                  // user-request domain
}
