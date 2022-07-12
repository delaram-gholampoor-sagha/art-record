package protocol

import "../libgo/protocol"

type StaffTakeLeave interface {
	StaffID() [16]byte         // staff domain
	Day() utc.DayElapsed       //
	DayHours() earth.DayHours  //
	Type() StaffTakeLeave_Type //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // user-request domain
}

type StaffTakeLeave_StorageServices interface {
	Save(st StaffTakeLeave) protocol.Error

	Count(staffID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (st StaffTakeLeave, err protocol.Error)

}

type StaffTakeLeave_Type uint8

const (
	StaffTakeLeave_Type_Unset StaffTakeLeave_Type = iota

	// usually four weeks of paid leave for each year of service. Shift-workers are entitled to five weeks of leave
	StaffTakeLeave_Type_Vacation
	StaffTakeLeave_Type_Travel
	StaffTakeLeave_Type_Rest
	StaffTakeLeave_Type_Study
	StaffTakeLeave_Type_Marriage
	StaffTakeLeave_Type_Casual

	StaffTakeLeave_Type_Sick // time off given by the company to allow employees to recover from an illness and take care of their health
	StaffTakeLeave_Type_Pain // Women period, ...
	StaffTakeLeave_Type_TemporaryDisability

	// Certain weather conditions make it extremely difficult, dangerous, and sometimes impossible for employees to travel to work.
	// Such conditions include heavy snow, a hurricane, or flash flooding.
	StaffTakeLeave_Type_AdverseWeather
	// or Time off in lieu (TOIL) for employees who have clocked in more hours than they were required to can be eligible for compensatory days off
	StaffTakeLeave_Type_Compensatory
	// Companies whose employees have served them for more than 3 years often
	// give these employees a sabbatical leave to reward them for their loyalty and hard work.
	StaffTakeLeave_Type_Sabbatical
	// Paid or unpaid time off gift
	StaffTakeLeave_Type_Gift
	// leave without pay
	StaffTakeLeave_Type_Unpaid

	StaffTakeLeave_Type_FamilyEvents
	// From taking care of the newborn to recovering from the delivery, maternity leave is an important time for new mothers.
	StaffTakeLeave_Type_Maternity
	// Paternity leave is granted to new fathers—  husbands or partners of a pregnant woman, surrogate parent,
	// or someone who adopted a child— to take care of their newborns without any worry.
	StaffTakeLeave_Type_Paternity
	// Losing a loved one is an unavoidable situation and in such events, employees take sudden leave.
	// Most HRs give their employees 3 to 7 days as bereavement leave, depending on the closeness of the relative.
	StaffTakeLeave_Type_Bereavement

	// Public holidays are days that are given as leave by the government.
	// Such holidays must be observed by every institution— schools, banks, government offices, and even private companies.
	StaffTakeLeave_Type_PublicHoliday
	// Christmas, Eid, Easter, Holi, Yom Kippur— your employee is sure to place importance on religious holidays that
	// they celebrate and would want the day off to spend time with their family and observe the festival.
	StaffTakeLeave_Type_ReligiousHolidays
	StaffTakeLeave_Type_JuryDuty
	StaffTakeLeave_Type_Protest
	StaffTakeLeave_Type_Volunteer
)


type (
	StaffTakeLeave_Service_Register_Request interface{
		StaffID() [16]byte         
   	Day() utc.DayElapsed       
   	DayHours() earth.DayHours 
	  Type() StaffTakeLeave_Type 
	}

	StaffTakeLeave_Service_Register_Response interface{
	  NumberOfVersion() protocol.NumberOfVersion
	}
	
	StaffTakeLeave_Service_Count_Request interface{
		StaffID() [16]byte
	}
	
	StaffTakeLeave_Service_Count_Response interface{
		NumberOfVersion() protocol.NumberOfVersion
	}
	StaffTakeLeave_Service_Get_Request interface{
		StaffID() [16]byte
		VersionOffset() uint64
	}
	
	StaffTakeLeave_Service_Get_Response interface{
		StaffTakeLeave
		NumberOfVersion() protocol.NumberOfVersion
	}
	
	StaffTakeLeave_Service_Last_Request interface{
		StaffID() [16]byte
	}
	
	StaffTakeLeave_Service_Last_Response interface{
		StaffTakeLeave
		NumberOfVersion() protocol.NumberOfVersion
	}
)