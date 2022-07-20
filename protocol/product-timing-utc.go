/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type ProductTimingUTC interface {
	ProductID() [16]byte      // product domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours // what hours in a day this service active to provide services
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type ProductTimingUTC_StorageServices interface {
	Save(pt ProductTimingUTC) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(productID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimingUTC, nv protocol.NumberOfVersion, err protocol.Error)
}


type (
	ProductTimingUTC_Service_Register_Request interface {
		ProductID() [16]byte      
	  Day() utc.DayElapsed      
	  DayHours() earth.DayHours
	}

	ProductTimingUTC_Service_Register_Response = protocol.NumberOfVersion

)

type (
	ProductTimingUTC_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductTimingUTC_Service_Count_Response = protocol.NumberOfVersion
	
)

type (
	ProductTimingUTC_Service_Get_Request interface {
		ProductID() [16]byte
		versionOffset() uint64
	}

	ProductTimingUTC_Service_Get_Response1 = ProductTimingUTC
	ProductTimingUTC_Service_Get_Response2 = protocol.NumberOfVersion
	
)