/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

// PosStatus indicate the domain record data fields.
type PosStatus interface {
	PosID() [16]byte     // pos domain
	Status() Pos_Status  //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type PosStatus_StorageServices interface {
	Save(ps PosStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.versionOffset) (ps PosStatus, nv protocol.NumberOfVersion, err protocol.Error)

	FilterByStatus(status Pos_Status, offset, limit uint64) (posIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	protocol.EventTarget
}

type Pos_Status Quiddity_Status

const (
	Pos_Status_CashLimitInactivated = Pos_Status(Quiddity_Status_FreeFlag << iota)
)



type (
	PosStatus_Service_Register_Request interface {
		PosID() [16]byte
		Status() Pos_Status
	}

	PosStatus_Service_Register_Response = protocol.NumberOfVersion

)

type (
	PosStatus_Service_Count_Request interface {
		PosID() [16]byte
	}

	PosStatus_Service_Count_Response = protocol.NumberOfVersion

)



type (
	PosStatus_Service_Get_Request interface {
		PosID() [16]byte    
		versionOffset() protocol.versionOffset
	}

	PosStatus_Service_Get_Response1 = PosStatus
	PosStatus_Service_Get_Response2 = protocol.NumberOfVersion

)



type (
	PosStatus_Service_FilterByStatus_Request interface{
		Status() Pos_Status
		Offset() uint64
		Limit() uint64
	}
	PosStatus_Service_FilterByStatus_Response1 interface{
		PosIDs() [][16]byte
	}

	PosStatus_Service_FilterByStatus_Response2 = protocol.NumberOfVersion
)

