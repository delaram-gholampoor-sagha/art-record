package protocol

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
	Save(pc ProductComplement) (numbers uint64, err protocol.Error)

	Count(productID [16]byte, priority uint64) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, priority uint64, versionOffset uint64) (pc ProductComplement, numbers uint64, err protocol.Error)

	ListProducts(complementID [16]byte, offset, limit uint64) (productIDs [16]byte, numbers uint64, err protocol.Error)
	ListPriorities(complementID [16]byte, productID [16]byte, offset, limit uint64) (priorities []uint64, numbers uint64, err protocol.Error)
}
type (
	ProductComplement_Service_Register_Request interface {
		ProductID() [16]byte 
		Priority() uint64       
		ComplementID() [16]byte 
	}
	ProductComplement_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	ProductComplement_Service_Count_Request interface {
		ProductID() [16]byte
		Priority() uint64
	}
	ProductComplement_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}

	ProductComplement_Service_Get_Request interface {
		ProductID() [16]byte
		Priority() uint64       
		VersionOffset() uint64
	}
	ProductComplement_Service_Get_Response interface {
		ProductComplement
		Nv() protocol.NumberOfVersion
	}


	ProductComplement_Service_ListProducts_Request interface {
		ComplementID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductComplement_Service_ListProducts_Response interface {
		ProductIDs() [][16]byte
		Nv() protocol.NumberOfVersion
	}


	ProductComplement_Service_ListPriorities_Request interface {
		ComplementID() [16]byte
		ProductID() [16]byte
		Offset() uint64
		Limit() uint64
	}
	ProductComplement_Service_ListPriorities_Response interface {
		Priorities() []uint64
		Nv() protocol.NumberOfVersion
	}
)