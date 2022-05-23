/* For license and copyright information please see LEGAL file in repository */

package protocol

import (
	"time"
)

// UserOtherAppCache cache all foreign fundamental user data
type UserOtherAppCache interface {
	UserOtherAppID() [16]byte // user-other-app domain
	UserID() [16]byte         // user-status domain
	Gender() string           // In ISO format
	PassportID() string       // If exist mean user have that country Nationality
	Surename() string         // Firstname
	Middlename() string       //
	Givenname() string        // Lastname
	Nickname() string         //
	Username() string         // it is also be user email address
	Phone() string            //
	TaxID() string            //
	Picture() [16]byte        // Object UUID
	Bio() string              //
	LocationID() string       //
	PostalCode() string       //
	DateOfBirth() time.Time   //
	PlaceOfBirth() string     //
	VerifiedDate() time.Time  // If !=0 user got famous blue badge(tick) like instagram, twitter,...
	Time() time.Time          // Save time
}
