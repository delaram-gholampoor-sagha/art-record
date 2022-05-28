package protocol

import (
	"../libgo/protocol"
)

// QuiddityIcon or quiddity emoji use in many domains as category, departments, ...
type QuiddityIcon interface {
	QuiddityID() [16]byte // quiddity domain
	Icon() [16]byte       // object domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}
