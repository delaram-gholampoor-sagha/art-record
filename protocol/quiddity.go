package protocol

import "../libgo/protocol"

// Quiddity indicate the domain record data fields.
type Quiddity interface {
	QuiddityID() [16]byte // Unique content ID
	DomainID() uint64     // Domain MediaTypeID e.g. product, content, ... domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type Quiddity_StorageServices interface {
	Save(q Quiddity) (err protocol.Error)

	Get(quiddityID [16]byte) (q Quiddity, err protocol.Error)

	// FindByDomainID(domainID [16]byte, offset, limit uint64) (quiddityIDs [][16]byte, numbers uint64, err protocol.Error)
}

type (
	Quiddity_Service_Register_Request interface {
		DomainID() [16]byte
	}
	
	Quiddity_Service_Register_Response interface {
		QuiddityID() [16]byte
		Nv() protocol.NumberOfVersion
	}
	
	Quiddity_Service_Get_Request interface {
		QuiddityID() [16]byte
	}
	Quiddity_Service_Get_Response interface {
		Quiddity
	}
	
)