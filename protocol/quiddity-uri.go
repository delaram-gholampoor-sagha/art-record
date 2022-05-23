

package protocol

import (
	"../libgo/protocol"
)

// https://en.wikipedia.org/wiki/Uniform_Resource_Identifier
// https://en.wikipedia.org/wiki/Uniform_Resource_Name
// https://en.wikipedia.org/wiki/Electronic_Product_Code
type QuiddityURI interface {
	URI() []byte          // Locale name in the Computer world.
	QuiddityID() [16]byte // quiddity domain
	Time() protocol.Time  // Save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityURI_StorageServices interface {
	Save(q QuiddityURI) (err protocol.Error)

	Count(uri []byte) (numbers uint64, err protocol.Error)
	Get(uri []byte, versionOffset uint64) (q QuiddityURI, err protocol.Error)
	Last(uri []byte) (q QuiddityURI, numbers uint64, err protocol.Error)

	FindByQuiddityID(quiddityID [16]byte, offset, limit uint64) (uris [][]byte, numbers uint64, err protocol.Error)
}
