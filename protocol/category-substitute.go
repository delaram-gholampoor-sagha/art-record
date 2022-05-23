/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type CategorySubstitute interface {
	CategoryID() [16]byte   // category domain
	Priotry() int32         // Substitute priotry
	SubstituteID() [16]byte // category domain
	Time() protocol.Time    // Save time
	RequestID() [16]byte    // user-request domain
}
