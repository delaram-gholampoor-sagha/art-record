/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

type CategoryItem interface {
	CategoryID() [16]byte // category domain
	Priority() uint64     //
	ItemID() [16]byte     // any domain usually product, content, ... domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type CategoryItem_StorageServices interface {
	Save(ci CategoryItem) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(categoryID [16]byte, priority uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(categoryID [16]byte, priority uint64, vo protocol.versionOffset) (ci CategoryItem, nv protocol.NumberOfVersion, err protocol.Error)

	ListCategories(itemID [16]byte, offset, limit uint64) (categoryIDs [16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListPriorities(itemID [16]byte, categoryID [16]byte, offset, limit uint64) (priorities []uint64, nv protocol.NumberOfVersion, err protocol.Error)
}

// services:
// Last(categoryID [16]byte, priority uint64) (ci CategoryItem, nv protocol.NumberOfVersion, err protocol.Error)
// ListItems(categoryID [16]byte, offset, limit uint64) (items [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)


type (
	CategoryItem_Service_Register_Request interface {
		CategoryID() [16]byte  
		Priority() uint64 
		ItemID() [16]byte 		
	}

	CategoryItem_Service_Register_Response = protocol.NumberOfVersion

)

type (
	CategoryItem_Service_Count_Request interface {
		CategoryID() [16]byte
		Priority() uint64 
	}

	CategoryItem_Service_Count_Response = protocol.NumberOfVersion
	
)



type (
	CategoryItem_Service_Get_Request interface {
		CategoryID() [16]byte
		Priority() uint64 
		versionOffset() protocol.versionOffset
	}

	CategoryItem_Service_Get_Response1 = CategoryItem
	CategoryItem_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	CategoryItem_Service_ListCategories_Request interface {
		ItemID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	CategoryItem_Service_ListCategories_Response1 interface {
		CategoryID() [16]byte
	}

	CategoryItem_Service_ListCategories_Response2 = protocol.NumberOfVersion
	
)


type (
	CategoryItem_Service_ListPriorities_Request interface {
		ItemID() [16]byte
		CategoryID() [16]byte
		Offset() uint64
		Limit() uint64
	}

	CategoryItem_Service_ListPriorities_Response1 interface {
		Priority() uint64 
	}

	CategoryItem_Service_ListPriorities_Response2 = protocol.NumberOfVersion
)

