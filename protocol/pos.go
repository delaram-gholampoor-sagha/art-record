package protocol

import (
	"../libgo/protocol"
)

// POS indicate the domain record data fields.
// POS or point-of-sale
type POS interface {
	PosID() [16]byte     // quiddity domain
	Type() POS_Type      //
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type POS_StorageServices interface {
	Save(pos POS) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.VersionOffset) (pos POS, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	POS_Service_Register_Request interface {
		PossID() [16]byte
		Type() POS_Type
	}
	POS_Service_Register_Response interface {
		NumberOfVersion()  protocol.NumberOfVersion
	}
	
	POS_Service_Count_Request interface {
		PosID() [16]byte
	}
	POS_Service_Count_Response interface {
		NumberOfVersion()  protocol.NumberOfVersion
	}
	
	POS_Service_Get_Request interface {
		PosID() [16]byte
		Vo() protocol.VersionOffset
	}
	POS_Service_Get_Response interface {
		POS
		NumberOfVersion()  protocol.NumberOfVersion
	
	}
)

type POS_Type uint64

// Common flags
const (
	POS_Type_Unset POS_Type = 0
	POS_Type_Cash  POS_Type = (1 << iota)
	POS_Type_Location
	POS_Type_Provider
	POS_Type_Product
	// POS_Type_
)
