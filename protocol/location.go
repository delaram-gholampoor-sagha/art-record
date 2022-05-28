package protocol

import (
	"../libgo/protocol"
)

// Warehouse, zone, ...
// QuiddityID use in Coordinate domain too to get physical location of the Location.
type Location interface {
	QuiddityID() [16]byte   // quiddity domain.
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}
