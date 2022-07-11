package protocol

import "../libgo/protocol"

type ProductTimingUTC interface {
	ProductID() [16]byte      // product domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours // what hours in a day this service active to provide services
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type ProductTimingUTC_StorageServices interface {
	Save(pt ProductTimingUTC) (numbers uint64, err protocol.Error)

	Count(productID [16]byte) (numbers uint64, err protocol.Error)
	Get(productID [16]byte, versionOffset uint64) (pt ProductTimingUTC, numbers uint64, err protocol.Error)
}

type (
	ProductTimingUTC_Service_Register_Request interface {
		ProductID() [16]byte      
	  Day() utc.DayElapsed      
	  DayHours() earth.DayHours
	}

	ProductTimingUTC_Service_Register_Response interface {
		Nv() protocol.NumberOfVersion
	}

	ProductTimingUTC_Service_Count_Request interface {
		ProductID() [16]byte
	}

	ProductTimingUTC_Service_Count_Response interface {
		Nv() protocol.NumberOfVersion
	}

	ProductTimingUTC_Service_Get_Request interface {
		ProductID() [16]byte
		VersionOffset() uint64
	}

	ProductTimingUTC_Service_Get_Response interface {
		ProductTimingUTC
		Nv() protocol.NumberOfVersion
	}
)