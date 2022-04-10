package protocol

import "time"

type Reservation struct {
	ID         uint64
	// QuiddityID [16]byte
	// UserID     uint64
	// EventID    uint64
	// PriceID    uint64
	RelatedID [16]byte
	Status    uint8

	// داده ی اکتشافی ؟
	Number_Of_Users uint64
	// different studios ?
	Location        string
	// maybe someone has to prepare sth in order to make the studio preparable for use ??
	// i think its the same as having a commnet on sth ?
	Special_Request string
	CheckIn_Date    time.Time
	CheckOut_Date   time.Time
	Created_at      time.Time
	Updated_at      time.Time
}

type ReservationServices interface {
	ChangeReservationState(ReservID uint64) (string, error)
	DeleteReservation(reservationID uint64) error
	UpdateReservation(isPartial bool, reservation Reservation) (Reservation, error)
	GetReservation(reservationID uint64) (Reservation, error)
	FindReservationsByEvent(reservID uint64) ([]Reservation, error)
	FindReservationByUser(userID uint64) (Reservation, error)
	FindReservationByUsers(userID uint64) (Reservation, error)
}
