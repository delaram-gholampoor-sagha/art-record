package protocol

type LogIn_Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register_Request struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone_number string `json:"phone_number"`
	Full_name    string `json:"full_name"`
}

type AuthUserModel struct {
	// User  UserModel   `json:"user"`
	Token string `json:"token"`
}

type User_Authentication interface { 
	LogInUser(email string, password string)
	RegisterUser(email string, phone_number string, full_name string, password string)
}
