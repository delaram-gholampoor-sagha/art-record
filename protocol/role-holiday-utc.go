package protocol


type RoleHolidayUTC interface {
	RoleID() [16]byte         // role domain
	Week() utc.WeekElapsed    //
	Weekdays() utc.Weekdays   // what days in weekday in utc time week this role active to provide services
	DayHours() earth.DayHours // what hours in a day this role active to provide services
	Time() protocol.Time      // save time
	RequestID() [16]byte      // user-request domain
}
