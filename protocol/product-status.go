package protocol

import "../libgo/protocol"

// ProductStatus indicate the domain record data fields.
type ProductStatus interface {
	ProductID() [16]byte    // product domain
	Status() Product_Status //
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductStatus_StorageServices interface {
	Save(ps ProductStatus) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (ps ProductStatus, numbers uint64, err protocol.Error)

	FilterByStatus(status Product_Status, offset, limit uint64) (productIDs [][16]byte, numbers uint64, err protocol.Error)

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

	ProductStatus_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	ProductStatus_Service_Count_Request interface {
		ProductID() [16]byte    
		
	}

	ProductStatus_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
		
	}

	Product_Status_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	Product_Status_Service_Get_Response interface {
		Status() Product_Status
		Nv() protocol.NumberOfVersion

	}

	Product_Status_Service_FilterByStatus_Request interface{
		Status() Product_Status
		Offset() uint64
		Limit() uint64
	}
	Product_Status_Service_FilterByStatus_Response interface{
		ProductIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}
)