package protocol

import "time"

type Event struct {
	ID uint64
	// CommentGroupID [16]byte
	// QuiddityID     [16]byte
	RelatedID      [16]byte
	CommentGroupID uint8
	// published_Date  ?
	Start_Date time.Time
	End_Date   time.Time
	Created_At time.Time
	Updated_At time.Time
}

type EventServices interface {
	RegisterEvent(event Event) (Event, error)
	UpdateEvent(isPartial bool, event Event) (Event, error)
	DeleteEvent(eventID uint64) error
	FindEventByTime(time string, eventID uint64) (Event, error)
	FindEventByLocation(location string, eventID uint64) (Event, error)
	GetTenLastEvents() ([]Event, error)
}
