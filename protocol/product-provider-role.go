package protocol

import (
	"../libgo/protocol"
)

// ProductProviderRole or product operator is to supply a particular service.
type ProductProviderRole interface {
	ProductID() [16]byte // product domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}

// related UserIDs by RoleID use to get Coordinate(domain)
