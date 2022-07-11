package protocol

import "../libgo/protocol"

// ProductTimingInterval indicate the domain record data fields.
// e.g. flight duration, doctor visit, carwash, ...
type ProductTimingInterval interface {
	ProductID() [16]byte         // product domain
	Minimum() protocol.Duration  // Minimum time duration for this service
	Expected() protocol.Duration // Expected time duration for this service
	Maximum() protocol.Duration  // Maximum time duration for this service
	Time() protocol.Time         // save time
	RequestID() [16]byte         // user-request domain
}

type ProductTimingInterval_StorageServices interface {
	Save(pt ProductTimingInterval) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimingInterval, nv protocol.NumberOfVersion, err protocol.Error)
}

type (
	ProductTimingInterval_Service_Register_Request interface {
			ProductID() [16]byte         
			Minimum() protocol.Duration  
			Expected() protocol.Duration 
			Maximum() protocol.Duration  
	}
	
	ProductTimingInterval_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductTimingInterval_Service_Count_Request interface {
		ProductID() [16]byte
	}
	
	ProductTimingInterval_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}
	
	ProductTimingInterval_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}
	
	ProductTimingInterval_Service_Get_Response interface {
		ProductTimingInterval
		Nv() protocol.NumberOfVersion
	}
	
)