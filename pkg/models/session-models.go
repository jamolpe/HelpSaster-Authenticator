package models

import (
	"time"

	guuid "github.com/google/uuid"
)

type (
	// Session : the session object
	Session struct {
		ID        guuid.UUID
		User      *User     `json:"User" bson:"user"`
		Token     string    `json:"Token" bson:"token"`
		CreatedAt time.Time `json:"CreatedAt" bson:"createdat"`
	}
)
