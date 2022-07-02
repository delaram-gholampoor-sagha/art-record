package protocol

// UserOtherAppUsername indicate the domain record data fields.
// it is also be user email address
type UserOtherAppUsername interface {
	UserID() [16]byte        // user domain
	OrgID() [16]byte         // user domain
	Username() string        //
	Status() UserName_Status //
	RequestID() [16]byte     // user-request domain
}

type UserOtherAppUsername_StorageServices interface {
}

/*
	Business Services
*/

type (
	UserOtherAppUsername_Service_Insert_Request interface {
		Email
	}

	GetEmail_Request interface {
		UserID() [16]byte
	}
	GetEmail_Response interface {
		Email
	}

	IsEmailVerified_Request interface {
		Email
	}
	IsEmailVerified_Response interface {
		Verified() bool
	}

	SendEmailVerification_Request interface {
		Email
	}

	IsEmailExist_Request interface {
		Email
	}
	IsEmailExist_Response interface {
		Exist() bool
	}

	GetEmailsHistoryPerUserID_Request interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}
	GetEmailsHistoryPerUserID_Response interface {
		Emails() []Email
	}
)
