/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductReserveQuantity indicate the domain record data fields.
type ProductReserveQuantity interface {
	ProductID() [16]byte // product domain
	Quantity() uint64    // fix max reserve number
	Percent() uint64     // percent of the inventory
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type ProductReserveQuantity_StorageServices interface {
	Save(pr ProductReserveQuantity) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pr ProductReserveQuantity, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductReserveQuantity_Service_Register_Request interface {
		ProductID() [16]byte
		Quantity() uint64
		Percent() uint64
	}
	ProductReserveQuantity_Service_Register_Response = protocol.NumberOfVersion

)


type (
	ProductReserveQuantity_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductReserveQuantity_Service_Count_Response = protocol.NumberOfVersion
	
)



type (
	ProductReserveQuantity_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	
	ProductReserveQuantity_Service_Get_Response1 = ProductReserveQuantity
	ProductReserveQuantity_Service_Get_Response2 = protocol.NumberOfVersion
	
)
