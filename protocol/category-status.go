/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

type CategoryStatus interface {
	CategoryID() [16]byte    // category domain
	Status() Category_Status //
	Time() protocol.Time     // save time
	RequestID() [16]byte     // user-request domain
}

type CategoryStatus_StorageServices interface {
	Save(gs CategoryStatus) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(categoryID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(categoryID [16]byte, vo protocol.versionOffset) (gs CategoryStatus, nv protocol.NumberOfVersion, err protocol.Error)

	FilterByStatus(status Category_Status, offset, limit uint64) (categoryIDs [][16]byte , nv protocol.NumberOfVersion , err protocol.Error)

	// protocol.EventTarget
}

type Category_Status Quiddity_Status

const (
// Category_Status_ = Category_Status(Quiddity_Status_FreeFlag << iota)
)


type (
		CategoryStatus_Service_Register_Request interface {
		CategoryID() [16]byte  
		Status() Category_Status
			
	}
	CategoryStatus_Service_Register_Response = protocol.NumberOfVersion

)

type (
	CategoryStatus_Service_Count_Request interface {
		CategoryID() [16]byte	
	}

	CategoryStatus_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	CategoryStatus_Service_Get_Request interface {
		CategoryID() [16]byte 
		versionOffset() protocol.versionOffset
	}

	CategoryStatus_Service_Get_Response1 = CategoryStatus
	CategoryStatus_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	CategoryStatus_Service_FilterByStatus_Request interface {
		Status() CategoryStatus
		Offset() uint64
		Limit() uint64
	}
	CategoryStatus_Service_FilterByStatus_Response1 interface {
		CategoryIDs() [][16]byte
	}

	CategoryStatus_Service_FilterByStatus_Response2 = protocol.NumberOfVersion
)

