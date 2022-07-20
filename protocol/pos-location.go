/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

// PosLocation indicate the domain record data fields.
type PosLocation interface {
	PosID() [16]byte              // pos domain
	BuildingLocationID() [16]byte // building-location domain
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type PosLocation_StorageServices interface {
	Save(ps PosProvider) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.versionOffset) (ps PosLocation, nv protocol.NumberOfVersion, err protocol.Error)
	
}

type (

	PosLocation_Service_Register_Request interface {
		PosID() [16]byte
		BuildingLocationID() [16]byte
	}

	PosLocation_Service_Register_Response = protocol.NumberOfVersion

)

type (
	PosLocation_Service_Count_Request interface {
		PosID() [16]byte
	}

	PosLocation_Service_Count_Response = protocol.NumberOfVersion

)


type (

	PosLocation_Service_Get_Request interface {
		PosID() [16]byte    
		versionOffset() protocol.versionOffset
	}

	PosLocation_Service_Get_Response1 = 	PosLocation
	PosLocation_Service_Get_Response2 = protocol.NumberOfVersion
	
)