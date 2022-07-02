package protocol

type UserOtherAppNumber interface {
	UserID() [16]byte        // user domain
	OrgID() [16]byte         // user domain
	PhoneNumber() string //
	Number() uint64      // must start with country code e.g. (00)98-912-345-6789
	Status() UserOtherAppNumber_Status
	RequestID() [16]byte // Request RecordID to prove how this object created in the chain of other objects.
}

type UserOtherAppNumber_StorageServices interface {
	Save(pn UserOtherAppNumber) (numbers uint64, err protocol.Error)

	Count(personID [16]byte) (numbers uint64, err protocol.Error)
	Get(personID [16]byte, versionOffset uint64) (pn UserOtherAppNumber, numbers uint64, err protocol.Error)

	FindByNumber(number uint64, offset, limit uint64) (personIDs [][16]byte, numbers uint64, err protocol.Error)
}

// UserOtherAppNumber_Status indicate UserOtherAppNumber record status
type UserOtherAppNumber_Status uint8

// UserOtherAppNumber status
const (
	UserOtherAppNumber_Status_Unset      UserOtherAppNumber_Status = 0
	UserOtherAppNumber_Status_Registered UserOtherAppNumber_Status = (1 << iota)
	UserOtherAppNumber_Status_Removed                        // Empty
	UserOtherAppNumber_Status_Hidden                         // in this case user can't login by number
	UserOtherAppNumber_Status_Approved
	UserOtherAppNumber_Status_Blocked
)

/*
	Business Services
*/

const UserOtherAppNumberServiceRegister = "urn:giti:person-number.sabz.city:service:register"

type (
	InsertPhoneRequest interface {
		Phone() string
	}

	GetPhoneRequest interface {
		UserID() [16]byte
	}

	GetPhoneResponse interface {
		Phone
	}

	IsPhoneVerifiedRequest interface {
		Phone
	}

	IsPhoneVerifiedResponse interface {
		Verified() bool
	}

	SendPhoneVerificationRequest interface {
		Phone
	}

	IsPhoneExistRequest interface {
		Phone
	}

	IsPhoneExistResponse interface {
		Exist() bool
	}

	GetPhonesHistoryPerUserIDRequest interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}

	GetPhonesHistoryPerUserIDResponse interface {
		Phones() []Phone
	}
)

/*
	Errors
*/
