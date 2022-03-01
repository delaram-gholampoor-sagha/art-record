package protocol

import "time"

type event_content struct {
	ID              uint64    `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	UserID          uint64    `json:"user_id"`
	ReservID        uint64    `json:"reserv_id"`
	Created_At      time.Time `json:"created_at"`
	Updated_At      time.Time `json:"updated_at"`
	Expiration_Time time.Time `json:"expiration_time"`
}

type Event interface {
	// Event
	RegisterEvent(event Event) (event_content, error)
	UpdateEvent(isPartial bool, event event_content) (event_content, error)
	DeleteEvent(eventID uint64) error
	FindEventByTime(time string, eventID uint64) (event_content, error)
	FindEventByLocation(location string, eventID uint64) (event_content, error)
	// get the 10 last events 
	GetEvents() ([]event_content, error)
}
