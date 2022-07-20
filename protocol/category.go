/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

// Category indicate the domain record data fields.
// Category or Topic
type Category interface {
	CategoryID() [16]byte // quiddity domain
	ParentID() [16]byte   // CategoryID, Exist if it isn't top category.
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type Category_StorageServices interface {
	Save(c Category) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(categoryID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(categoryID [16]byte, vo protocol.versionOffset) (gs Category, nv protocol.NumberOfVersion, err protocol.Error)

	GetIDs(offset, limit uint64) (categoryIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
		Category_Service_Register_Request interface {
			ParentID() [16]byte
		}

		Category_Service_Register_Response1 interface {
				CategoryID() [16]byte
		}	

		Category_Service_Register_Response2 = protocol.NumberOfVersion

)

type (
		Category_Service_Count_Request interface {
				CategoryID() [16]byte

		}
		Category_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	Category_Service_Get_Request interface {
				CategoryID() [16]byte    
				versionOffset() protocol.versionOffset
	}

	Category_Service_Get_Response1 = 	Category
	Category_Service_Get_Response2 =  protocol.NumberOfVersion
	
)



type (	
		Category_Service_GetIDs_Request interface {
			Offset() uint64
			Limit() uint64
		}
		Category_Service_GetIDs_Response1 interface {
			CategoryIDs() [][16]byte
		}

		Category_Service_GetIDs_Response2 = protocol.NumberOfVersion
)