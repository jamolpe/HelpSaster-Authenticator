package sessioncore

import "authorization-service/pkg/models"

// SessionRepository : expected methods for session repository
type SessionRepository interface {
	SaveSession(session models.Session) error
	GetSessionByUserID(UserID string) (*models.Session, error)
}
