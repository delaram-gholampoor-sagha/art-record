package protocol


type PersonNumber interface {
	PersonID() [16]byte // user-status domain
	Number() uint64     // must start with country code e.g. (00)98-912-345-6789
	Status() PersonNumber_Status
	RequestID() [16]byte // Request RecordID to prove how this object created in the chain of other objects.
}

type PersonNumber_StorageServices interface {
	Save(pn PersonNumber) (err protocol.Error)

	Count(personID [16]byte) (numbers uint64, err protocol.Error)
	Get(personID [16]byte, versionOffset uint64) (pn PersonNumber, err protocol.Error)
	Last(personID [16]byte) (pn PersonNumber, numbers uint64, err protocol.Error)

	FindByNumber(number uint64, offset, limit uint64) (personIDs [][16]byte, numbers uint64, err protocol.Error)
}

// PersonNumber_Status indicate PersonNumber record status
type PersonNumber_Status uint8

// PersonNumber status
const (
	PersonNumber_Status_Unset      PersonNumber_Status = 0
	PersonNumber_Status_Registered PersonNumber_Status = (1 << iota)
	PersonNumber_Status_Removed                        // Empty
	PersonNumber_Status_Hidden                         // in this case user can't login by number
	PersonNumber_Status_Approved
	PersonNumber_Status_Blocked
)

/*
	Business Services
*/

const PersonNumberServiceRegister = "urn:giti:person-number.sabz.city:service:register"

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
