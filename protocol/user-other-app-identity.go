package protocol

// UserOtherAppIdentity cache all foreign fundamental user data
type UserOtherAppIdentity interface {
	UserID() [16]byte            // user domain
	OrgID() [16]byte             // user domain
	Gender() string              // In ISO format
	PassportID() string          // If exist mean user have that country Nationality
	Surname() string             // FirstName
	MiddleName() string          //
	GivenName() string           // LastName
	Nickname() string            //
	TaxID() string               //
	Picture() string             // picture url
	Bio() string                 //
	Location() string            //
	PostalCode() string          //
	DateOfBirth() protocol.Time  //
	PlaceOfBirth() string        //
	VerifiedDate() protocol.Time // If !=0 user got famous blue badge(tick) like instagram, twitter,...
	Time() protocol.Time         // save time
}
