package core

import "authorization-service/pkg/models"

// Repository : dtb repository needs the implementation of: UserRepository: methods SaveUser GetUserByEmail SessionRepository: methods SaveSession GetSessionByUserID
type Repository interface {
	UserRepository
	SessionRepository
}

// UserRepository : expected methods for user repository
type UserRepository interface {
	SaveUser(user models.User) error
	GetUserByEmail(user models.User) (models.User, error)
}

// SessionRepository : expected methods for session repository
type SessionRepository interface {
	SaveSession(session models.Session) error
	GetSessionByUserID(UserID string) (*models.Session, error)
}
