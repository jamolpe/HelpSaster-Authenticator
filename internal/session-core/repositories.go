package sessioncore

import (
	"go-sessioner/pkg/models"
)

// SessionRepository : expected methods for session repository
type SessionRepository interface {
	SaveSession(session models.Session) error
	GetSessionByUserID(userID string) (*models.Session, error)
	GetSessionByID(ID string) (*models.Session, error)
	UpdateSession(session models.Session) error
}
