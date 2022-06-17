package protocol

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

	Count(staffID [16]byte) (numbers uint64, err protocol.Error)
	Get(staffID [16]byte, versionOffset uint64) (st StaffTimingUTC, err protocol.Error)
	Last(staffID [16]byte) (st StaffTimingUTC, numbers uint64, err protocol.Error)
}
