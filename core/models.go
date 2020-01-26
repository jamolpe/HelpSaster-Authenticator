package core

// User : user model
type User struct {
	Email, Name string
	Password    string `json:"-"`
	Age         uint8
}

// LoginUserResponse : response from login
type LoginUserResponse struct {
	Logged bool
	User   User
	Token  string
}

// LoginUserData : request body from login
type LoginUserData struct {
	Email    string
	Password string
}
