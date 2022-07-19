/* For license and copyright information please see LEGAL file in repository */

package art

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
	Save(q QuiddityURI) (nv protocol.NumberOfVersion,  err protocol.Error)

	Count(uri []byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(uri []byte, versionOffset uint64) (q QuiddityURI, nv protocol.NumberOfVersion  , err protocol.Error)


	FindByQuiddityID(quiddityID [16]byte, offset, limit uint64) (uris [][]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	QuiddityURI_Service_Register_Request interface {
		URI() []byte
		QuiddityID() [16]byte
	}

	QuiddityURI_Service_Register_Response = protocol.NumberOfVersion

)

type (
	QuiddityURI_Service_Count_Request interface {
		URI() []byte	
	}

	QuiddityURI_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	QuiddityURI_Service_Get_Request interface {
		URI() []byte
		VersionOffset() uint64 
	}

	QuiddityURI_Service_Get_Response1 = QuiddityURI
	QuiddityURI_Service_Get_Response2 = protocol.NumberOfVersion
	
)


type (
	QuiddityURI_Service_FindQuiddityID_Request interface {
		URI() []byte
		QuiddityID() [16]byte
	}	

	QuiddityURI_Service_FindQuiddityID_Response1 = protocol.NumberOfVersion
)