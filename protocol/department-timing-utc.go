package protocol


type DepartmentTimingUTC interface {
	DepartmentID() [16]byte   // department domain
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this department active to provide services
	DayHours() earth.DayHours // what hours in day in utc time week this department active to provide services
	Time() protocol.Time      // Save time
	RequestID() [16]byte      // user-request domain
}
