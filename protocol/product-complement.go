/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductComplement indicate the domain record data fields.
type ProductComplement interface {
	ProductID() [16]byte    // product domain
	Priority() uint64       // Complement priority
	ComplementID() [16]byte // product domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductComplement_StorageServices interface {
	Save(pc ProductComplement) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte, priority uint64) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, priority uint64, versionOffset uint64) (pc ProductComplement, nv protocol.NumberOfVersion, err protocol.Error)

	ListProducts(complementID [16]byte, offset, limit uint64) (productIDs [16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	ListPriorities(complementID [16]byte, productID [16]byte, offset, limit uint64) (priorities []uint64, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductComplement_Service_Register_Request interface {
		ProductID() [16]byte 
		Priority() uint64       
		ComplementID() [16]byte 
	}

	ProductComplement_Service_Register_Response = protocol.NumberOfVersion
)

type (
	ProductComplement_Service_Count_Request interface {
		ProductID() [16]byte
		Priority() uint64
	}

	ProductComplement_Service_Count_Response = protocol.NumberOfVersion
)

type (
	ProductComplement_Service_Get_Request interface {
		ProductID() [16]byte
		Priority() uint64       
		versionOffset() uint64
	}

	ProductComplement_Service_Get_Response1 = ProductComplement
	ProductComplement_Service_Get_Response2 = protocol.NumberOfVersion

)


type (
	ProductComplement_Service_ListProducts_Request interface {
		ComplementID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductComplement_Service_ListProducts_Response1 interface {
		ProductIDs() [][16]byte
	}

	ProductComplement_Service_ListProducts_Response2 = protocol.NumberOfVersion
	
)



type (
	ProductComplement_Service_ListPriorities_Request interface {
		ComplementID() [16]byte
		ProductID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductComplement_Service_ListPriorities_Response1 interface {
		Priorities() []uint64
	}

	ProductComplement_Service_ListPriorities_Response2 = protocol.NumberOfVersion
)