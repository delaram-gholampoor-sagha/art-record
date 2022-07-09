package protocol

type RoleTimingUTC interface {
	RoleID() [16]byte         // role domain
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this role active to provide services
	DayHours() earth.DayHours // what hours in day in utc time week this role active to provide services
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}

type RoleTimingUTC_StorageServices interface {
	Save(rt RoleTimingUTC) protocol.Error

	Count(roleID [16]byte) (numbers uint64, err protocol.Error)
	Get(roleID [16]byte, versionOffset uint64) (rt RoleTimingUTC, err protocol.Error)
	
}

type (
	RoleTimingUTC_Service_Register_Request interface{
		RoleID() [16]byte        
  	Weekdays() utc.Weekdays  
  	DayHours() earth.DayHours 
	}

	RoleTimingUTC_Service_Register_Response interface{
		Numbers() uint64
	}
	
	RoleTimingUTC_Service_Count_Request interface{
		RoleID() [16]byte
	}
	
	RoleTimingUTC_Service_Count_Response interface{
		Numbers() uint64
	}
	RoleTimingUTC_Service_Get_Request interface{
		RoleID() [16]byte
		VersionOffset() uint64
	}
	
	RoleTimingUTC_Service_Get_Response interface{
		RoleTimingUTC
		Numbers() uint64
	}
	
)