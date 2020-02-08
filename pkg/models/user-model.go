package models

type (
	// LoginUser : login model to identify
	LoginUser struct {
		Email    string
		Password string
	}
	// User : user model
	User struct {
		Email, Name string
		Password    string `json:"-"`
		Age         uint8
	}

	// SessionUser : user info session
	SessionUser struct {
		Logged bool
		User   *User
		Token  string
	}
)
