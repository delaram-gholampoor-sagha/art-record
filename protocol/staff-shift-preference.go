package protocol

import "../libgo/protocol"

// StaffShift indicate the domain record data fields.
type StaffShift interface {
	StaffID() [16]byte   // staff domain
	Day() utc.DayElapsed //
	ShiftID() [16]byte   // org-shift domain
	Time() protocol.Time // save time
	RequestID() [16]byte // user-request domain
}



type StaffShift_StorageServices interface {
	Save(ss StaffShift) protocol.Error

	Count(staffID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (ss StaffShift, err protocol.Error)
	
}


type (
	
	StaffShift_Service_Register_Request interface{
	 StaffID() [16]byte   // staff domain
	 Day() utc.DayElapsed //
	 ShiftID() [16]byte   // org-shift domain
	}

	StaffShift_Service_Register_Response interface{
   	NumberOfVersion() protocol.NumberOfVersion
	}

	StaffShift_Service_Count_Request interface{
		StaffID() [16]byte
	}

	StaffShift_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
	StaffShift_Service_Get_Request interface{
		StaffID() [16]byte
		VersionOffset() uint64
	}

	StaffShift_Service_Get_Response interface{
		StaffShift
	}

)