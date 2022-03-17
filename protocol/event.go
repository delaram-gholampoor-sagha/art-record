package protocol

import "time"

type event_content struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// pic ?
	Pic_URL             string    `json:"pic_url"`
	Expiration_DateTime time.Time `json:"expiration_datetime"`
	Issue_datetime      string    `json:"issue_datetime"`
	Created_At          time.Time `json:"created_at"`
	Updated_At          time.Time `json:"updated_at"`
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
