package protocol

import (
	"../libgo/protocol"
)

type VoucherValue interface {
	VoucherID() [16]byte           // voucher domain
	Each() uint8                   // Each time use
	Percentage() uint64            // Max discount percentage per use
	Price() protocol.AmountOfMoney // Max discount price per use
	Time() protocol.Time           // Save time
	RequestID() [16]byte           // user-request domain
}


