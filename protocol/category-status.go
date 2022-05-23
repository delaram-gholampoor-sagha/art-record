/* For license and copyright information please see LEGAL file in repository */

package org

import (
	"../libgo/protocol"
)

type CategoryStatus interface {
	CategoryID() [16]byte    // category domain
	Status() Category_Status //
	Time() protocol.Time     // Save time
	RequestID() [16]byte     // user-request domain
}

type CategoryStatus_StorageServices interface {
}

type Category_Status uint8
