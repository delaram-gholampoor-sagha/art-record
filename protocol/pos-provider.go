/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

// PosProvider indicate the domain record data fields.
// Provider or operator is to limit pos to specific staffs.
type PosProvider interface {
	PosID() [16]byte     // pos domain
	RoleID() [16]byte    // role domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type PosProvider_StorageServices interface {
	Save(ps PosProvider) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.versionOffset) (ps PosProvider, nv protocol.NumberOfVersion, err protocol.Error)
	
}


type (
	PosProvider_Service_Register_Request interface {
		PosID() [16]byte
		RoleID() [16]byte
	}
	PosProvider_Service_Register_Response = protocol.NumberOfVersion

)

type (

	PosProvider_Service_Count_Request interface {
		PosID() [16]byte
	}

	PosProvider_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	PosProvider_Service_Get_Request interface {
		PosID() [16]byte    
		versionOffset() protocol.versionOffset
	}

	PosProvider_Service_Get_Response1 =  PosProvider
	PosProvider_Service_Get_Response2 =  protocol.NumberOfVersion

)

	