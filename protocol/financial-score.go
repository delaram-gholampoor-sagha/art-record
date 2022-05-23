

package protocol

import (
	"../libgo/protocol"
)

type FinancialScore interface {
	UserID() [16]byte    // user-status domain
	Score() int32        //
	Time() protocol.Time // Save time
	RequestID() [16]byte // user-request domain
}
