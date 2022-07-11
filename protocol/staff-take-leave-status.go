package protocol

import "../libgo/protocol"

type StaffTakeLeaveStatus interface {
	StaffID() [16]byte             // staff domain
	Day() utc.DayElapsed           //
	Status() StaffTakeLeave_Status //
	Time() protocol.Time           // save time
	RequestID() [16]byte           // user-request domain
}

type StaffTakeLeaveStatus_StorageServices interface {
	Save(ss StaffTakeLeaveStatus) protocol.Error

	Count(staffID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (ss StaffTakeLeaveStatus, err protocol.Error)
	

	// FilterByStatus(status StaffTakeLeave_Status, offset, limit uint64) (staffIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// protocol.EventTarget
}

type StaffTakeLeave_Status Quiddity_Status

const (
	StaffTakeLeave_Status_NeedAlternate = StaffTakeLeave_Status(Quiddity_Status_FreeFlag << iota)
	StaffTakeLeave_Status_NeedAlternateApprove
	StaffTakeLeave_Status_AlternateApprove
	StaffTakeLeave_Status_NeedManagerApprove
	StaffTakeLeave_Status_ManagerApprove
)



type (
	StaffTakeLeaveStatus_Service_Register_Request interface{
		StaffID() [16]byte             
	  Day() utc.DayElapsed          
	  Status() StaffTakeLeave_Status 
	}

		StaffTakeLeaveStatus_Service_Register_Response interface{
	   Nv() protocol.NumberOfVersion
	}
	
	
	StaffTakeLeaveStatus_Service_Count_Request interface{
		StaffID() [16]byte
	}
	
	StaffTakeLeaveStatus_Service_Count_Response interface{
		Nv() protocol.NumberOfVersion
	}
	StaffTakeLeaveStatus_Service_Get_Request interface{
		StaffID() [16]byte
		VersionOffset() uint64
	}
	
	StaffTakeLeaveStatus_Service_Get_Response interface{
		StaffTakeLeaveStatus
	}
	
)