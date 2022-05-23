

package protocol

import "time"



type PictureRating interface {
	ObjectID() [16]byte     //
	Rating() picture.Rating //
	Time() time.Time        // Save time
	RequestID() [16]byte    // user-request domain
}
