/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductMultipleUse indicate the domain record data fields.
type ProductMultipleUse interface {
	ProductID() [16]byte        // product domain
	Maximum() uint64            // Maximum time this service can be served to buyer
	Timeout() protocol.Duration // from first use
	Time() protocol.Time        // save time
	RequestID() [16]byte        // user-request domain
}

type ProductMultipleUse_StorageServices interface {
	Save(pm ProductMultipleUse) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pm ProductMultipleUse, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductMultipleUse_Service_Register_Request interface {
		ProductID() [16]byte        
		Maximum() uint64           
		Timeout() protocol.Duration 
	}

	ProductMultipleUse_Service_Register_Response = protocol.NumberOfVersion

)



type (
	ProductMultipleUse_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductMultipleUse_Service_Count_Response = protocol.NumberOfVersion
	
)


type (
	ProductMultipleUse_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	ProductMultipleUse_Service_Get_Response1 interface {
		ProductMultipleUse
		NumberOfVersion() protocol.NumberOfVersion
	}

	ProductMultipleUse_Service_Get_Response2 = protocol.NumberOfVersion
)