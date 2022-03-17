package protocol

import "time"

type Reservation_Content struct {
	ID              uint64    `json:"id"`
	Text            string    `json:"text"`
	UserID          string    `json:"user_id"`
	// cancelled , confirmed , on hold
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


// all the days are disabled untill the admin has set the date and time that the studio is available.


// Studio is available for course-related use by students currently enrolled in a Cinematic Arts or Communications Studies course.




// restaurant ??
// Pre-ordering option

// https://tyrocity.com/topic/reservation-process-procedure/
// hotels
// Reservation is a complete process of booking that is conducted by to parties i.e. one guest or customer and next one is hotel reservation section.
// The procedure of reservation are:
// A)     Enquiry for reservations:
// First step of reservation is registered as enquiry of reservation where reservation personnel conduct 
// a brief question answer section with guest or customer to gain various knowledge about reservation which he/she wants to make.
// The point that are to be recorded are note down below:

// i)      Name of the guest

// ii)     Date of arrival

// iii)    Date of departure

// iv)     Desired room type

// v)      Desired room rate

// vi)     Required number of rooms

// vii)    Desired room plan

// viii)   Number of pare.

// ix)     Contact address and number (details)

// x)      Special recommendation


// B)     Determining the room availability:
// The second most important step in reservation is determining the room availability.
// In this process we check the demand of guest encoded during the first step.
// The availability can be checked by referring to forecast chart, conventional chart or density chart. 
// In fully automated system we can begin check in the same availability by computerized system or software.

// C) Excepting or dying the request:
// Now after check in the room availability in hotel we are able to expect or deny the reservation request. 
// Expecting can be done if the request is fully validated by the availability formats and system, where as vice versa can be conducted for denying.
// Exceptional: In few cases reservation personnel an up sale another room of same criteria or higher by conveying or motivate the customer or guest.
// D) Documenting the reservation details:
// After accepting or denying reservation next process is documentation.
// Documentation is conducted by reservation staff by completing undone part of reservation form with the necessary details of guest who is about to stay.
// E) Confirming the reservation request:
// After documenting the details of a guest we need to confirm the room to a guest confirmation are generally done 
// by dispatching a confirmation letter to the guest by having a guaranteed reservation and may be sometime non-guarantee too.
// F) Maintaining the reservation record:
// After confirming a room to the guest we need to make a reservation record for each and every booking made. 
// The records can be maintained or processed in two ways:

// a) Documenting the original reservation: In this process we file and record the original reservation details 
// and if we are working in a computerized system we will be recording a printed form of reservation done.
// The documentation are done in a basis of date of arrival and afterward in assurance to the surname of guest.
// b) Modification of reservation: The next step in processing is changing of the details that are recorded in reservation form.
// In this case, we need to attach the different ammessdement/ correction form or slip with the original one.

// G) Compiling the reservation report:
// Now after completion of reservation record we need to prepare the report of reservation dept. on either basis of date, week, month or year.


// 6 Stages Of Reservation Process in Hotel Industry
// 1. Identification of the sources of business:
//  The hotel normally identifies two major sources of business: Free Independent Travellers (FIT) and associated Travellers (Groups).
// The FIT component is simply a non-group affiliate seeking overnight accommodations.
// Regardless of whether a request is made by an individual, a family or a collection of friends, 
// a request may still be termed as FIT reservation if it meets certain qualifications.
// 2. Communication with the Hotel:
// There are numerous ways a hotel can learn about the guest’s request for available space. 
// The potential guest can connect with the hotel via a telephone, email, online, Social Media,
// other devices or referral systems may all become important, communication links between the sources of business and the hotel property.


// 3. Formulation of the reservation request:
// Once a hotel is contacted by a guest about reserving space its staff must intelligently formulate 
// the reservation request and complete a reservation record. The reservation record is the permanent 
// file form used to store all the information pertinent to a specific reservation or in current era creating 
// a computerised reservation record in the hotel management software or PMS.

// In formulating the reservation request, the three most important factors are a) the expected date of arrival and departure b) 
// the desired room type c) the number of rooms needed. Using this information the hotel staff can search the room availability file via the designated arrival date. If the date is not closed the reservation agent proceeds by checking room type and the number of rooms available.

// Along with this preliminary data, a reservation record will usually be judged complete when it includes the following information:
// Guest’s Profile Data.
// Guest’s group/company affiliation.
// Arrival date and time and departure date.
// The number of nights.
// Number of Pax.
// Room type requested.
// The number of rooms required.
// Type of bed.
// Rate quoted / Rate Code.
// Date reservation received
// Reservation Status.
// Billing Instruction
// Pickup Details.
// Any other Details or special remarks.
// Source Code
// Market Code

// 5. Confirmation notification:
// The acceptance of the reservation is followed by a notification confirming the blocking of a room or group of rooms.
// The reservation confirmation is sent to the booker or guest by email, fax, SMS, social media or online chat (Instant messenger).

// The confirmation establishes :
// A check on the entered reservation request data and satisfaction of the reservation request.
// An agreement of room rates.
// An agreement of method of payment.
// A statement of the hotel’s cancellation procedure.
// 6. Modification and pre-arrival activity:
// Once the reservation has been accepted and confirmed many changes to reservation record can occur prior to the guest’s arrival.
// Changes in the arrival or departure dates, number of rooms required, adding pickup details, reconfirmation, updating deposit status, number of pax and outright cancellations are examples of common pre-arrival activities. Any time a change in a reservation record is required, the reservation record must be pulled and updated according to the guest request to keep the reservation-related data up to date.

// Also searching and retrieving the correct reservation record for any modification and cancellation request is crucial to an effective reservation process.
// Also. while processing the cancellation below details need to be updated on the reservation record for any future reference.
// Caller Name.
// Caller contact details.
// Reason for Cancellation.
// Cancellation reference number.
// Date of Cancellation (Recorded automatically by the hotel software or enter manually).
// User or Reservation Agent Name Who processed the cancellation (Recorded automatically by the hotel software or enter manually).

















// https://www.tutorialspoint.com/front_office_management/front_office_management_reservation.htm
// Reservation of the hotel accommodation is one of the important responsibilities of the front office department. A potential guest contacts a hotel for availability of the desired type of accommodation and any allied services that the hotel offers. The front office department needs to react to the enquiry of the guests.

// For a guest, reservation increases the chances of a better deal for assured accommodation on arrival. For a hotel, reservation can enable a better management of guest experience during usual as well as peak seasons. Reservation procedure varies depending on the size and brand of the hotel and the reservation system employed.

// Let us know the details how the front office handles reservations.

// Types of Hotel Reservation Systems
// An efficient and effective reservation system is what adds to the hotel’s profitability. Following are the most popular reservation systems −

// Whitney System of Reservation
// It was developed in 1940 by Whitney Paper Corporation from New York, hence the name. This is a conventional manual reservation system the hotels used to follow during pre-computer days in the hotels. It contains the following setup for reservation −

// Slip for request of accommodation reservation
// Whitney slip that records guest name, accommodation type, number, and duration of stay
// Temporary/Permanent arrival slip
// Guest bill
// Guest registration card
// Correspondence file
// Bedroom journal that records daily occupancy of the guest with date, guest name, room type, and room number
// Let us see how a Whitney slip and the bedroom journal looks like.

// Whitney Slip
// Guest Name	Date of Arrival	Room Type	Room Rate	Date of Departure
// Mode of Reservation	Reserved By	Date Received
// Reservation Agency if Any:
// Billing Instructions	Date of Confirmation
// Bedroom Journal
// Date: ___________
// Guest Name	Room Number	Room Particulars
// Though this system proved efficient, it generated a lot of paperwork with occasional scope for errors. The drawbacks were overcome by the central reservation system.

// Central Reservation System (CRS)
// It is a computerized reservation system that reduces paperwork and can handle large amount of reservation data effortlessly.

// In this system, since the guest data and reservation data are stored on the storage disks of the computers, it can be accessed at wish. It is stored in the form of a database of collection of records which can enable searching, adding, removing, or updating any guest related data.

// The computerized reservation system not only helps to make guest reservations but also helps to forecast how many accommodations can be reserved in an upcoming time period.

// This is how a CRS typically works −

// CRS
// The guests of hotel sales agents call for checking room availability. It is forwarded to the front office reservation staff. The staff finds out details about the requirement and checks the availability of desired accommodation in the database. According to the reservation policies and procedures, the reservation staff member then notifies or suggests the reception about the accommodation availability and takes further appropriate action.

// The Role of Internet in Reservation
// The Internet has brought a momentum in the hospitality business as well. It facilitates seamless management of a hotel’s offices located at various places and their various departments.

// The hotel businesses are actively working on the Internet 24 hours a day, seven days a week. The Internet has simplified complex system of reservations. It enables Online Hotel Management Systems (OHMS) such as Hotelogix to help guests reserve accommodation of their choice fast and conveniently. The guests of the hotel can access rate charts, accommodation availability, check-in and check-out timings, details about the restaurants, and so on, at their own convenience.

// Sources of Reservations
// People travel for various reasons such as personal as well as for MICE. There are various sources from whom the requests of reservation pour in −

// Direct Request from Guests − The prospective guests can approach individually to the hotel for reservation of accommodation mostly when they are single travelers or family travelers.

// Request from Travel Agent − They can approach the hotel for booking accommodations for group travelers.

// Request from Corporate Agent − An organization can request a hotel to reserve accommodations for their employees, clients, or visitors.

// Request from Airlines − The airlines can reserve accommodations for their working staff for routine stay as well as in case of flight cancellations.

// Request from Institutions − Various SMERF or NGO institutions request to reserve hotels for sports people, delegations of embassies, or performing-art program groups, workshop groups, and alike who travel to different location.

// Managing Reservations
// The first step in reserving an accommodation is to check if the requested kind of accommodation is available for selling for a specific period of time. It is done by checking forecast boards or computerized systems.

// Accepting Reservation of Accommodation
// Reservation of an accommodation is accepted if the desired type of accommodation is available in the hotel for selling. If it is not available during a rush season or if the guest is in urgent need, the staff member suggests for almost similar alternative accommodation by stating its amenities and facilities.

// Reservation is accepted in the following cases in conjunction with the availability of the accommodation −

// Is the guest new to the hotel?
// Does the guest have good credentials with the hotel regarding payment and behavior?
// Is the guest a VIP?
// Denying Reservation of Accommodation
// Denial of reservation directly means loss of revenue. But there are certain situations when the reservation staff turns down the reservation for the guests or agents. The potential causes of denying reservation are −

// All accommodations in hotel booked − In such case, the reservation staff refuses the reservation politely and suggests an alternative hotel in the same area or different property of the same owner in a nearby area.

// Requested type of accommodation not available − In such case, the reservation staff suggests an alternate accommodation.

// Guest/Agent blacklisted − Some guests or agents are blacklisted due to their history of payment dues against the hotel. In such case, the reservation clerk seeks for reservation manager’s advice.

// Finally, the reservation section of the front office prepares the list of the reservations for the day and sends it to the front desk. The list also contains vital information such as if the guest is new or repeat, guest preferences about room location or décor. The rooms are then prepared by housekeeping.

// Cancelling a Reservation of Accommodation
// This is yet another event when the hotel loses business with a guest. Though the fact is overt loss of revenue, the front office staff must react to it politely and gracefully. The staff member also needs to convey any cancellation charges the guest must pay while cancelling the reservation. Cancellation is done in the following steps −

// Finding out details of the guest and respective reserved accommodation.
// Verifying charges of cancellation, if any.
// Notifying the guest about cancellation charges.
// Cancelling the reservation in the system.
// Updating the system for accommodation availability.
// Confirming the guest about the cancellation.
// Generating Reservation Report
// Reservation reports are generated for the sake of helping the management find trends and making forecast about business.

// The reports typically are of the following types −

// Occupancy report
// Special arrival report
// Revenue forecast report
// Turnaway report

