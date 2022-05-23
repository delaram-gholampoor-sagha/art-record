

package protocol

import (
	"../libgo/protocol"
)

// https://en.wikipedia.org/wiki/Expiration_date
type ProductBestBefore interface {
	ProductID() [16]byte         // product-status domain
	Duration() protocol.Duration // from production time
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}