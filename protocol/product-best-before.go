/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

// ProductBestBefore indicate the domain record data fields.
// https://en.wikipedia.org/wiki/Expiration_date
type ProductBestBefore interface {
	ProductID() [16]byte         // product domain
	Duration() protocol.Duration // from production time
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductBestBefore_StorageServices interface {
	Save(pb ProductBestBefore) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pb ProductBestBefore, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductBestBefore_Service_Register_Request interface {
		ProductID() [16]byte 
		Duration() protocol.Duration     
	}

	ProductBestBefore_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductBestBefore_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductBestBefore_Service_Count_Response = protocol.NumberOfVersion
	
)


type (	
	ProductBestBefore_Service_Get_Request interface {
		ProductID() [16]byte    
		VersionOffset() uint64
	}
	
	ProductBestBefore_Service_Get_Response1 =	ProductBestBefore
	ProductBestBefore_Service_Get_Response2 = protocol.NumberOfVersion
	
)
