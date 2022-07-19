/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// UserVehicle
// Vehicle is a conveyance that transports people or objects
type UserVehicle interface {
	UserID() [16]byte    // user domain
	ThingID() [16]byte   //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}
