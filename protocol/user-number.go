package protocol

type UserNumber interface {
	UserID() protocol.UserID // user domain
	OrgID() protocol.UserID  // user domain. TenantID
	Number() string          // PhoneNumber
	// Number() uint64          // without country code e.g. {(00)98-}912-345-6789
	Status() UserNumber_Status //
	Time() protocol.Time       // save time
	RequestID() [16]byte       // Request RecordID to prove how this object created in the chain of other objects.
}

type UserNumber_StorageServices interface {
	Save(pn UserNumber) (nv protocol.NumberOfVersion, err protocol.Error)

	Count(userID, orgID [16]byte) (nv protocol.NumberOfVersion, err protocol.Error)
	Get(userID, orgID [16]byte, vo protocol.VersionOffset) (pn UserNumber, nv protocol.NumberOfVersion, err protocol.Error)

	FindByNumber(number uint64, offset, limit uint64) (userIDs [][16]byte, nv protocol.NumberOfVersion, err protocol.Error)
}

// UserNumber_Status indicate UserNumber record status
type UserNumber_Status uint8

// UserNumber status
const (
	UserNumber_Status_Unset      UserNumber_Status = 0
	UserNumber_Status_Registered UserNumber_Status = (1 << iota)
	UserNumber_Status_Removed                      // Empty
	UserNumber_Status_Hidden                       // in this case user can't login by number
	UserNumber_Status_Approved
	UserNumber_Status_Blocked
)

/*
	Business Services
*/

type (
	UserNumber_Service_Register_Request interface {
		Phone() string
	}
)

type (
	UserNumber_Service_Get_Request interface {
		UserID() protocol.UserID
	}
	UserNumber_Service_Get_Response interface {
		Phone
	}
)

type (
	IsPhoneVerifiedRequest interface {
		Phone
	}
	IsPhoneVerifiedResponse interface {
		Verified() bool
	}
)

type (
	SendPhoneVerificationRequest interface {
		Phone
	}
)

type (
	IsPhoneExistRequest interface {
		Phone
	}
	IsPhoneExistResponse interface {
		Exist() bool
	}
)

type (
	GetPhonesHistoryPerUserIDRequest interface {
		UserID() protocol.UserID
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}
	GetPhonesHistoryPerUserIDResponse interface {
		Phones() []Phone
	}
)
