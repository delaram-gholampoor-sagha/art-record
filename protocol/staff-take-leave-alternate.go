package protocol

import "../libgo/protocol"

// StaffTakeLeaveAlternate or staff replace
type StaffTakeLeaveAlternate interface {
	StaffID() [16]byte     // staff domain
	Day() utc.DayElapsed   //
	PresentedID() [16]byte // staff domain
	Time() protocol.Time   // save time
	RequestID() [16]byte   // user-request domain
}


type StaffTakeLeaveAlternate_StorageServices interface {
	Save(stla StaffTakeLeaveAlternate) protocol.Error

	Count(staffID [16]byte) (numbers uint64, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (stla StaffTakeLeaveAlternate, err protocol.Error)
	

	// FilterByStatus(status StaffTakeLeave_Status, offset, limit uint64) (staffIDs [][16]byte, numbers uint64, err protocol.Error)
	// protocol.EventTarget
}


type (
	StaffTakeLeaveAlternate_Service_Register_Request interface{
		StaffID() [16]byte     
   	Day() utc.DayElapsed  
  	PresentedID() [16]byte 
	}

		StaffTakeLeaveAlternate_Service_Register_Response interface{
	   Nv() protocol.NumberOfVersion
	}
	
	
	StaffTakeLeaveAlternate_Service_Count_Request interface{
		StaffID() [16]byte
	}
	
	StaffTakeLeaveAlternate_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	StaffTakeLeaveAlternate_Service_Get_Request interface{
		StaffID() [16]byte
		VersionOffset() uint64
	}
	
	StaffTakeLeaveAlternate_Service_Get_Response interface{
		StaffTakeLeaveAlternate
	}
	
)