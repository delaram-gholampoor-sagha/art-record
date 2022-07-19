/* For license and copyright information please see LEGAL file in repository */

package art

import "../libgo/protocol"

type StaffOvertimeStatus interface {
	StaffID() [16]byte            // staff domain
	Day() utc.DayElapsed          //
	Status() StaffOvertime_Status //
	Time() protocol.Time          // save time
	RequestID() [16]byte          // user-request domain
}

type StaffOvertimeStatus_StorageServices interface {
	Save(gs StaffOvertimeStatus) protocol.Error

	Count(staffID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (gs StaffOvertimeStatus, nv protocol.NumberOfVersion , err protocol.Error)
	
	// FilterByStatus(status StaffOvertime_Status, offset, limit uint64) (staffIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
	// protocol.EventTarget
}

type StaffOvertime_Status Quiddity_Status

const (
	StaffOvertime_Status_NeedManagerApprove = StaffOvertime_Status(Quiddity_Status_FreeFlag << iota)
	StaffOvertime_Status_ManagerApprove
)


type (
	StaffOvertimeStatus_Service_Register_Request interface{
		StaffID() [16]byte            // staff domain
  	Day() utc.DayElapsed          //
  	Status() StaffOvertime_Status //
	}

	StaffOvertimeStatus_Service_Register_Response = protocol.NumberOfVersion

)

type (
 StaffOvertimeStatus_Service_Count_Request interface{
		StaffID() [16]byte
	}
	
	StaffOvertimeStatus_Service_Count_Response = protocol.NumberOfVersion

)



type (
	StaffOvertimeStatus_Service_Get_Request interface{
		StaffID() [16]byte
		VersionOffset() uint64
	}
	
	StaffOvertimeStatus_Service_Get_Response1 = StaffOvertimeStatus
	StaffOvertimeStatus_Service_Get_Response2 = protocol.NumberOfVersion

)



