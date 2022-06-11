package protocol


type DepartmentHolidayUTC interface {
	DepartmentID() [16]byte   // department domain
	Week() utc.WeekElapsed    //
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this role active to provide services
	DayHours() earth.DayHours // what hours in a day this role active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
