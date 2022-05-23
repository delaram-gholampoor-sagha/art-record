package protocol

import "time"

type Reservation interface {
	ID() [16]byte
	UserID() [16]byte
	EventID() [16]byte
	RequestID() [16]byte
	InvoiceID() [16]byte
	Status() Reservation_Status
	CheckIn_Date() time.Time
	CheckOut_Date() time.Time
}


type Reservation_StorageServices interface {
	Save(r Reservation) error

	Count(reservationID [16]byte) (numbers uint64, err error)
	Get(reservationID [16]byte, versionOffset uint64) (r Reservation, err error)
	Last(reservationID [16]byte) (r Reservation, numbers uint64, err error)

	GetReservationValidityDate(id [16]byte) (startDate string, endDate string, err error)
	GetReservationExpirationDate(id [16]byte) (startDate string, endDate string, err error)
	GetReservationStatus(id [16]byte) (status uint8, err error)

	FindReservationsByEventID(eventID [16]byte) (resrv Reservation, err error)
	ListUsersByEventID(eventID [16]byte) (id [][16]byte, err error)
	FindReservationStatusByInvoiceID(invoiceID [16]byte) (status uint8, err error)
}

type Reservation_Status uint8 


const (
	Reservation_Status_Unset  Reservation_Status = iota
	Reservation_Status_onHold
	Reservation_Status_Confirmed
	Reservation_Status_Cancelled 
	Reservation_Status_Approved
)


// the domain events help you to express , explicitly ,
// the domain rules.
// is there any side effects that i need to be aware of ?
// are the domain events synchronous or asynchronous ?
// domain evens / integration events : notifications about
// sth that just happened
// domain event => ?? : messages pushed to a domain-event-dispatcher
// integration event => propogate comittted transaction updates to
// additional subsystems (other microservices , bounded contnext)
// they should occure only if the entity is successfully  persisted,
// otherwise as if the entire operation never happened

// must be asynchronous operation between multiple microservices (other bouned context)
// alternatively you can have the aggregate root subscribed for events
// raised by the member of its aggregate (child entity)
// CommandHandler/EventHandler : commandhandler should be processed
// only once

// an event is sth that happened in the past , it should not change
// therefore it must be an immutable class ... read_only properties
// there is only one-way to update the object , you can only set values
// when you created.

// you have to decide => dispatching the domain events right before or right after committing
// a transaction ?
// it helps us determines whether you will include the side effects
// as part of the same transaction or in different transaction = eventual consistency
// we want the side effect happens in the same logical transaction
// but not neccessarily in the same scope of raising the domain event
// just before commiting the transaction we dispatch our events to their
// respective handlers

// storing domain events in additional database tables ? so they can be part of
// an original transaction
//
