

package protocol

import (
	"../libgo/picture"
	"../libgo/protocol"
)

type PictureRating interface {
	ObjectID() [16]byte     //
	Rating() picture.Rating //
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}
