/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

type CategoryComplement interface {
	CategoryID() [16]byte   // category domain
	Priority() uint64       // Complement priority
	ComplementID() [16]byte // category domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type CategoryComplement_StorageServices interface {
	Save(cs CategoryComplement) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(categoryID [16]byte, priority uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(categoryID [16]byte, priority uint64, vo protocol.versionOffset) (cs CategoryComplement, nv protocol.NumberOfVersion, err protocol.Error)

	ListCategories(complementID [16]byte, offset, limit uint64) (categoryIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListPriorities(complementID, categoryID [16]byte, offset, limit uint64) (priorities []uint64, nv protocol.NumberOfVersion, err protocol.Error)
}

// services:
// ListSubstitute(categoryID [16]byte, offset, limit uint64) (complementIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)


type (
	CategoryComplement_Service_Register_Request interface {
		CategoryID() [16]byte  
		Priority() uint64 
		ComplementID() [16]byte 		
	}

	CategoryComplement_Service_Register_Response = protocol.NumberOfVersion

)

type (
	CategoryComplement_Service_Count_Request interface {
		CategoryID() [16]byte
		Priority() uint64 
	}

	CategoryComplement_Service_Count_Response = protocol.NumberOfVersion
)

type (
	CategoryComplement_Service_Get_Request interface {
		CategoryID() [16]byte
		Priority() uint64 
		versionOffset() protocol.versionOffset
	}

	CategoryComplement_Service_Get_Response1 = CategoryComplement
	CategoryComplement_Service_Get_Response2 = protocol.NumberOfVersion
)

type (

	CategoryComplement_Service_ListCategories_Request interface {
		ComplementID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	CategoryComplement_Service_ListCategories_Response1 interface {
		CategoryIDs() [][16]byte			
	}

	CategoryComplement_Service_ListCategories_Response2 = protocol.NumberOfVersion
)

type (
	CategoryComplement_Service_ListPriorities_Request interface {
		ComplementID() [16]byte
		CategoryID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	CategoryComplement_Service_ListPriorities_Response1 interface {
		priorities() uint64 			
	}

	CategoryComplement_Service_ListPriorities_Response2 = protocol.NumberOfVersion
)

