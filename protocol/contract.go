package protocol

import (
	"../libgo/protocol"
)

type Contract interface {
	ContractID() [16]byte    // contract domain
	TemplateID() [16]byte    // contract-template domain
	Data() map[string]string //
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain
}
