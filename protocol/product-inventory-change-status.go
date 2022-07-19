/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductInventoryChangeStatus indicate the domain record data fields.
type ProductInventoryChangeStatus interface {
	ProductInventoryChangeID() [16]byte    // product-inventory-change domain
	Status() ProductInventoryChange_Status // overall
	Time() protocol.Time                   // save time
	RequestID() [16]byte                   // user-request domain
}

type ProductInventoryChangeStatus_StorageServices interface {
	Save(ps ProductInventoryChangeStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productInventoryChangeID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productInventoryChangeID [16]byte, versionOffset uint64) (ps ProductInventoryChangeStatus, nv protocol.NumberOfVersion, err protocol.Error)

	FilterByStatus(status ProductInventoryChange_Status, offset, limit uint64) (productInventoryChangeIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)

	protocol.EventTarget
}

type ProductInventoryChange_Status Quiddity_Status

// DC abbreviation for distribution-center or Warehouse
const (
	ProductInventoryChange_Status_OutcomeDCRequested = ProductInventoryChange_Status(Quiddity_Status_FreeFlag << iota)
	ProductInventoryChange_Status_IncomeDCRequested
	ProductInventoryChange_Status_OutcomeDCApproved
	ProductInventoryChange_Status_IncomeDCApproved
	ProductInventoryChange_Status_Void
)


type (
		ProductInventoryChangeStatus_Service_Register_Request interface {
		ProductInventoryChangeID() [16]byte    
		Status() ProductInventoryChange_Status
	}
	ProductInventoryChangeStatus_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductInventoryChangeStatus_Service_Count_Request interface {
		ProductInventoryChangeID() [16]byte
	}

	ProductInventoryChangeStatus_Service_Count_Response = protocol.NumberOfVersion
)


type (
	ProductInventoryChangeStatus_Service_Get_Request interface {
		productInventoryChangeID() [16]byte
		VersionOffset() uint64
	}

	ProductInventoryChangeStatus_Service_Get_Response1 = 	ProductInventoryChangeStatus
	ProductInventoryChangeStatus_Service_Get_Response2 = protocol.NumberOfVersion
	
)

type (
	ProductInventoryChangeStatus_Service_FilterByStatus_Request interface {
		Status() ProductInventoryChange_Status
		Offset()
		Limit() uint64
	}
	
	ProductInventoryChangeStatus_Service_FilterByStatus_Response1 interface {
		ProductInventoryChangeIDs() [][16]byte
	}

	ProductInventoryChangeStatus_Service_FilterByStatus_Response2 = protocol.NumberOfVersion
)


