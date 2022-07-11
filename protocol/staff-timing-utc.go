package protocol

import "../libgo/protocol"

type StaffTimingUTC interface {
	StaffID() [16]byte        // staff domain
	ShiftID() [16]byte        // org-shift domain
	Day() utc.DayElapsed      //
	DayHours() earth.DayHours // what hours in a day this staff must active to provide services or do projects
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type StaffTimingUTC_StorageServices interface {
	Save(st StaffTimingUTC) protocol.Error

	Count(staffID [16]byte) (numbers uint64, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (st StaffTimingUTC, err protocol.Error)
	
}


type (
	StaffTimingUTC_Service_Register_Request interface{
			StaffID() [16]byte        
			ShiftID() [16]byte        
			Day() utc.DayElapsed      
			DayHours() earth.DayHours 
	}

	StaffTimingUTC_Service_Register_Response interface{
	  Nv() protocol.NumberOfVersion
	}
	
	StaffTimingUTC_Service_Count_Request interface{
		StaffID() [16]byte
	}
	
	StaffTimingUTC_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	StaffTimingUTC_Service_Get_Request interface{
		StaffID() [16]byte
		VersionOffset() uint64
	}
	
	StaffTimingUTC_Service_Get_Response interface{
		StaffTimingUTC
		Nv() protocol.NumberOfVersion
	}
	
	StaffTimingUTC_Service_Last_Request interface{
		StaffID() [16]byte
	}
	
	StaffTimingUTC_Service_Last_Response interface{
		StaffTimingUTC
		Nv() protocol.NumberOfVersion
	}
)