package protocol

import "../libgo/protocol"

// QuiddityURI indicate the domain record data fields.
// https://en.wikipedia.org/wiki/Uniform_Resource_Identifier
// https://en.wikipedia.org/wiki/Uniform_Resource_Name
// https://en.wikipedia.org/wiki/Electronic_Product_Code
type QuiddityURI interface {
	URI() []byte          // Locale name in the Computer world.
	QuiddityID() [16]byte // quiddity domain
	Time() protocol.Time  // save time
	RequestID() [16]byte  // user-request domain
}

type QuiddityURI_StorageServices interface {
	Save(q QuiddityURI) (err protocol.Error)

	Count(uri []byte) (numbers uint64, err protocol.Error)
	Get(uri []byte, versionOffset uint64) (q QuiddityURI, err protocol.Error)


	FindByQuiddityID(quiddityID [16]byte, offset, limit uint64) (uris [][]byte, numbers uint64, err protocol.Error)
}

type (
	
	QuiddityURI_Service_Register_Request interface {
		URI() []byte
		QuiddityID() [16]byte
	}

	QuiddityURI_Service_Register_Response interface {
		Numbers() uint64
	}
	QuiddityURI_Service_Count_Request interface {
		URI() []byte	
	}

	QuiddityURI_Service_Count_Response interface {
		Numbers() uint64
	}

	QuiddityURI_Service_Get_Request interface {
		URI() []byte
		VersionOffset() uint64 
	}

	QuiddityURI_Service_Get_Response interface {
		QuiddityURI
	}


	QuiddityURI_Service_FindQuiddityID_Request interface {
		URI() []byte
		QuiddityID() [16]byte
	}	

	QuiddityURI_Service_FindQuiddityID_Response interface {
		Numbers() uint64
	}
)