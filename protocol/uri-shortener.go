package protocol

import (
	"../libgo/protocol"
)

type URIShortener interface {
	ID() string                  // Short Unique ID
	URI() string                 // UUID of picture object.
	UserID() [16]byte           // user-status domain
	Status() URIShortener_Status //
	Time() protocol.Time         // Save time
	RequestID() [16]byte         // user-request domain
}

const shortenerPath = "/s"
