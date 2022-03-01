package protocol

import "time"

type Reservation_Content struct {
	ID              uint64    `json:"id"`
	Text            string    `json:"text"`
	UserID          string    `json:"user_id"`
	State           string    `json:"state"`
	EventID         string    `json:"event_id"`
	Product_PriceID string    `json:"product_price_id"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
}

type Reservation interface {
	RegisterReservation(reservation Reservation_Content) (Reservation_Content, error)
	// ??
	ChangeReservationState(ReservID uint64) (string, error)
	DeleteReservation(reservationID uint64) error
	// permission from admin ?
	// UpdateReservation(isPartial bool, reservation Reservation_Content) (Reservation_Content, error)
	GetReservation(reservationID uint64) (Reservation_Content, error)
	// FindReservationsByEvent(reservID uint64) ([]Reservation_Content, error)
	// FindReservationByUser(userID uint64) (Reservation_Content , error)
	
}
