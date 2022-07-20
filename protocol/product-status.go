/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductStatus indicate the domain record data fields.
type ProductStatus interface {
	ProductID() [16]byte    // product domain
	Status() Product_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductStatus_StorageServices interface {
	Save(ps ProductStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (ps ProductStatus, nv protocol.NumberOfVersion, err protocol.Error)

	FilterByStatus(status Product_Status, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	protocol.EventTarget
}

type Product_Status Quiddity_Status

const (
// Product_Status_ = Product_Status(Quiddity_Status_FreeFlag << iota)
)


type (
	ProductStatus_Service_Register_Request interface {
		ProductID() [16]byte    
		Status() Product_Status 
	}

	ProductStatus_Service_Register_Response = protocol.NumberOfVersion
)

type (
	ProductStatus_Service_Count_Request interface {
		ProductID() [16]byte    
	}

	ProductStatus_Service_Count_Response = protocol.NumberOfVersion
)


type (
	Product_Status_Service_Get_Request interface {
		ProductID() [16]byte
		versionOffset() uint64
	}
	Product_Status_Service_Get_Response1 interface {
		Status() Product_Status
	}

	Product_Status_Service_Get_Response2 = protocol.NumberOfVersion
	
)

type (
	Product_Status_Service_FilterByStatus_Request interface{
		Status() Product_Status
		Offset() uint64
		Limit() uint64
	}
	Product_Status_Service_FilterByStatus_Response1 interface{
		ProductIDs() [][16]byte
	}

	Product_Status_Service_FilterByStatus_Response2 = protocol.NumberOfVersion
)