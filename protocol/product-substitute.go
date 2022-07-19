/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductSubstitute indicate the domain record data fields.
type ProductSubstitute interface {
	ProductID() [16]byte    // product domain
	Priority() uint64       // Substitute priority
	SubstituteID() [16]byte // product domain
	Time() protocol.Time    // save time
	RequestID() [16]byte    // user-request domain
}

type ProductSubstitute_StorageServices interface {
	Save(ps ProductSubstitute) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (ps ProductSubstitute, nv protocol.NumberOfVersion, err protocol.Error)

	FindBySubstitute(substituteID [16]byte, offset, limit uint64) (productIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductSubstitute_Service_Register_Request interface{
		ProductID() [16]byte    
		Priority() uint64      
		SubstituteID() [16]byte 
	}
	ProductSubstitute_Service_Register_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}

)

type (
		ProductSubstitute_Service_Count_Request interface{
		ProductID() [16]byte
	}

	ProductSubstitute_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (
		ProductSubstitute_Service_Get_Request interface{
		ProductID() [16]byte
		VersionOffset() uint64
	}

	ProductSubstitute_Service_Get_Response interface{
		ProductSubstitute
		NumberOfVersion() protocol.NumberOfVersion
	}
	
)


type (

	ProductSubstitute_Service_FilterByType_Request interface{
		SubstituteID() uint64
		Offset() uint64
		Limit() uint64
	}

	ProductSubstitute_Service_FilterByType_Response interface{
		ProductIDs() [][16]byte
		NumberOfVersion() protocol.NumberOfVersion
	}
)