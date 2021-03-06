/* For license and copyright information please see LEGAL file in repository */

package art

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

	Count(staffID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (st StaffTimingUTC, err protocol.Error)
	
}


type (
	StaffTimingUTC_Service_Register_Request interface{
			StaffID() [16]byte        
			ShiftID() [16]byte        
			Day() utc.DayElapsed      
			DayHours() earth.DayHours 
	}

	StaffTimingUTC_Service_Register_Response = protocol.NumberOfVersion
)

type (
	StaffTimingUTC_Service_Count_Request interface{
		StaffID() [16]byte
	}
	
	StaffTimingUTC_Service_Count_Response = protocol.NumberOfVersion

)


type (
	StaffTimingUTC_Service_Get_Request interface{
		StaffID() [16]byte
		versionOffset() uint64
	}
	
	StaffTimingUTC_Service_Get_Response1 = StaffTimingUTC
	StaffTimingUTC_Service_Get_Response2 = protocol.NumberOfVersion
)

