package models

type (
	// LoginUser : login model to identify
	LoginUser struct {
		Email    string
		Password string
	}
	// RegisterUser : register model
	RegisterUser struct {
		Email    string
		Password string
		Name     string
		LastName string
		Age      uint8
	}
	// User : user model
	User struct {
		Email, Name string
		Password    string `json:"-"`
		Age         uint8
	}

	// AuthUser : user info session
	AuthUser struct {
		Logged bool   `json:"Logged"`
		User   *User  `json:"User"`
		Token  string `json:"Token"`
	}
)
