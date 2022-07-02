package protocol


import (
	"../libgo/protocol"
)

type ContractAnalytic interface {
	CoordinateID() [16]byte // contract domain

	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
