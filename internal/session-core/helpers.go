package sessioncore

import (
	"go-sessioner/pkg/models"
	"time"
)

func mapSessionFromAuthUser(authUser *models.AuthUser) *models.Session {
	session := new(models.Session)
	session.Token = authUser.Token
	session.UserID = authUser.User.Email
	session.User = authUser.User
	session.CreatedAt = time.Now()
	return session
}
