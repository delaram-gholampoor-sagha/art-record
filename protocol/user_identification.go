package protocol

import "time"

type User struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	// HashPassword []byte    `json:"hashpassword"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Avatar    string    `json:"avatar"`
}

// this is the inforamtion that you are going to retrieve when it is a public request
type PublicUser struct {
	Id int64 `json:"id"`
	// Full_name  string `json:"full_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

// this is the inforamtion that you are going to retrieve when it is an internal request
type PrivateUser struct {
	Id         int64  `json:"id"`
	Full_Name  string `json:"full_name"`
	Email      string `json:"email"`
	Created_at string `json:"date_created"`
	Status     string `json:"status"`
	// Password    string `json:"password"`
}

type User_Identification interface {
	// Identification
	GetUser(userID uint64) (User, error)
	UpdateUser(isPartial bool, user User) (User, error)
	// disable or delete ?
	DisableUser(userID uint64) error
	
	FindUserByEmail(email string) ([]User , error)
	FindUserByFullName(fullname string) ([]User, error)
	FindUserByReservedEvent(eventID uint64 ) ([]User ,error) 
}
