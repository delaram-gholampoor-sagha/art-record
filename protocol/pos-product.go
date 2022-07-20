/* For license and copyright information please see LEGAL file in repository */

package art

import (
	"../libgo/protocol"
)

// PosServiceArea indicate the domain record data fields.
// Save shortcut to regular products to easily add to invoice
type PosProduct interface {
	PosID() [16]byte     // pos domain
	ProductID() [16]byte // product domain.
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}

type PosProduct_StorageServices interface {
	Save(pp PosProduct) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(posID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(posID [16]byte, vo protocol.versionOffset) (pp PosProduct,nv protocol.NumberOfVersion ,  err protocol.Error)


}

type (
		PosProduct_Service_Register_Request interface {
		PosID() [16]byte
		ProductID() [16]byte
	}

	PosProduct_Service_Register_Response = protocol.NumberOfVersion

)

type (
	PosProduct_Service_Count_Request interface {
		PosID() [16]byte
	}

	PosProduct_Service_Count_Response = protocol.NumberOfVersion

)



type (
	PosProduct_Service_Get_Request interface {
		PosID() [16]byte    
		versionOffset() protocol.versionOffset
	}

	PosProduct_Service_Get_Response1 = 	PosProduct
	PosProduct_Service_Get_Response2 = protocol.NumberOfVersion
	
)

