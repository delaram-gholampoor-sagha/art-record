package protocol

import "time"

/*
	Storage Data Structure
*/

type UserEmail interface {
	UserID() [16]byte        // user-status domain
	Email() string           //
	Status() UserEmailStatus //
	RequestID() [16]byte     // user-request domain
}

type UserEmailStatus uint8

const (
	UserEmailStatus_Unset UserEmailStatus = iota
	UserEmailStatus_Active
	UserEmailStatus_Inactive
	UserEmailStatus_Blocked
	UserEmailStatus_Reserved
)

// EmailService is interface of Email storage layer
type EmailService interface {
	Create(email Email) error
	Find(email string) (Email, error)
	Last(userID [16]byte) (Email, error)
	Lasts(userID [16]byte, count int) ([]Email, error)
	Meantime(userID [16]byte, start, end time.Time) ([]Email, error)
	CountVersionUntil(userID [16]byte, until time.Time) (int, error)
}

/*
	Business Services
*/

type (
	InsertEmailRequest interface {
		Email
	}

	GetEmailRequest interface {
		UserID() [16]byte
	}

	GetEmailResponse interface {
		Email
	}

	IsEmailVerifiedRequest interface {
		Email
	}

	IsEmailVerifiedResponse interface {
		Verified() bool
	}

	SendEmailVerificationRequest interface {
		Email
	}

	IsEmailExistRequest interface {
		Email
	}

	IsEmailExistResponse interface {
		Exist() bool
	}

	GetEmailsHistoryPerUserIDRequest interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}

	GetEmailsHistoryPerUserIDResponse interface {
		Emails() []Email
	}
)
