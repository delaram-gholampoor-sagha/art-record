package protocol

import "time"

type User struct {
	ID           uint
	FullName     string
	Email        string
	PhoneNumber  string
	Status       string
	// authorization
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Avatar       string
}

type User_Identification interface {
	GetUser(userID uint64) (User, error)
	UpdateUser(isPartial bool, user User) (User, error)
	DisableUser(userID uint64) error

	FindUserByEmail(email string) ([]User, error)
	FindUserByFullName(fullname string) ([]User, error)
	FindUserByReservedEvent(eventID uint64) ([]User, error)

	ResetPassword(email string) (string, error)
	ChangePassword(email string , password string) (string , error)
	SendPasswordResetEmail(email string) (string , error)
	ResendApproveSMSCode(PhoneNumber string) (string , error) 
	LogOutUser(userId uint64) error
}
