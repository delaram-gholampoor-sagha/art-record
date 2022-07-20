/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type Building interface {
	BuildingID() [16]byte // quiddity domain
	AreaID() [16]byte     // area domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type Building_StorageServices interface {
	Save(c Building) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(BuildingID [16]byte, AreaID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(BuildingID [16]byte, AreaID [16]byte, vo protocol.versionOffset) (gs Building, nv protocol.NumberOfVersion, err protocol.Error)

	GetIDs(offset, limit uint64) (BuildingIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	Building_Service_Register_Request interface {
		BuildingID() [16]byte
		AreaID() [16]byte
	}

	Building_Service_Register_Response = protocol.NumberOfVersion

)

type (
	Building_Service_Count_Request interface {
		BuildingID() [16]byte
		AreaID() [16]byte
	}
	
	Building_Service_Count_Response = protocol.NumberOfVersion

)

type (
	Building_Service_Get_Request interface {
		BuildingID() [16]byte
		AreaID() [16]byte
		versionOffset() protocol.versionOffset
	}

	Building_Service_Get_Response1 = Building
	Building_Service_Get_Response2 =  protocol.NumberOfVersion
)

type (
	Building_Service_GetIDs_Request interface {
		Offset() uint64
		Limit() uint64
	}
	Building_Service_GetIDs_Response1 interface {
		BuildingIDs() [][16]byte
	}

	Building_Service_GetIDs_Response2 = protocol.NumberOfVersion
)
