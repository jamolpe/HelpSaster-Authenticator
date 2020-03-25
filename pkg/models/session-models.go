package models

import (
	"time"
)

type (
	// Session : the session object
	Session struct {
		ID        string
		User      *User     `json:"User"`
		Token     string    `json:"Token"`
		UserID    string    `json:"UserID"`
		CreatedAt time.Time `json:"CreatedAt"`
	}
)
