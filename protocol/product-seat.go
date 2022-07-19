/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductSeat indicate the domain record data fields.
type ProductSeat interface {
	ProductID() [16]byte    // product domain
	SeatCapacity() uint64   // max seat number
	Type() ProductSeat_Type //
	MapID() [16]byte        // object domain. to show background and rows and seats. suggest use SVG format
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductSeat_StorageServices interface {
	Save(ps ProductSeat) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (ps ProductSeat, nv protocol.NumberOfVersion, err protocol.Error)
}

type ProductSeat_Type uint8

const (
	ProductSeat_Type_Unset           ProductSeat_Type = 0
	ProductSeat_Type_ChooseSeatByOrg ProductSeat_Type = (1 << iota)
	ProductSeat_Type_ChooseSeatByAttendanceTime
)


type (
	ProductSeat_Services_Register_Request interface{
		ProductID() [16]byte
		SeatCapacity() uint64
		Type() ProductSeat_Type
		MapID() [16]byte 
	}
	ProductSeat_Services_Register_Response = protocol.NumberOfVersion

)

type (
		ProductSeat_Service_Count_Request interface{
		ProductID() [16]byte
	}
	ProductSeat_Service_Count_Response = protocol.NumberOfVersion

)



type (
	ProductSeat_Service_Get_Request interface{
		ProductID() [16]byte
		VersionOffset() uint64
	}
	
	ProductSeat_Service_Get_Response1 = ProductSeat
	ProductSeat_Service_Get_Response2 = protocol.NumberOfVersion
)

