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
		Email    string `bson:"email"`
		Name     string `bson:"name"`
		Password string `json:"-" bson:"password"`
		Age      uint8  `bson:"age"`
	}

	// AuthUser : user info login response
	AuthUser struct {
		Logged bool   `json:"Logged"`
		User   *User  `json:"User"`
		Token  string `json:"Token"`
	}
)
