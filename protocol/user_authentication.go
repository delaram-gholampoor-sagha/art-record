package protocol

type Authentication struct {
	UserID   uint64
	Password [32]byte
}

type AuthenticationServices interface {
	LogInUser(email string, password string) error
	RegisterUser(email string, phone_number string, full_name string, password string) error
}
