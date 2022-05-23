package protocol

import "time"


type Event interface {
	ID() [16]byte
	QuiddityID() [16]byte
	QuiddityPriceID() [16]byte
	Location() string
	Start_Date() time.Time
	End_Date() time.Time
}

type EventServices_StorageServices interface {
	Save(e Event) error

	Count(eventID [16]byte) (length uint64, err error)
	Get(eventID [16]byte, versionOffset uint64) (c Event, err error)
	Last(eventID [16]byte) (c Event, length uint64, err error)

	FindEventByQuiddityID(quiddityID [16]byte) (id [16]byte, err error)
	FindEventByLocation(location string) (ev Event, err error)

	// reserv
	ListUserGroups(userID [16]byte, offset, limit uint64) (ids [][16]byte, numbers uint64, err error)
	ListEventsByDuration(duration string) (ev []Event, err error)
}
