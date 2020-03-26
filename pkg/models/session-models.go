package models

import (
	"time"
)

type (
	// Session : the session object
	Session struct {
		ID        string
		User      *User     `json:"User" bson:"user"`
		Token     string    `json:"Token" bson:"token"`
		UserID    string    `json:"UserID" bson:"userid"`
		CreatedAt time.Time `json:"CreatedAt" bson:"createdat"`
	}
)
