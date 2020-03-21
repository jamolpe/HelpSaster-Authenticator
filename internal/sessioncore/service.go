package sessioncore

import (
	"authorization-service/pkg/models"

	mgerror "authorization-service/pkg/errors"

	gologger "github.com/jamolpe/go-logger"
)

// SessionInterface : session core service
type SessionInterface interface {
	GetSession(userID string) (*models.Session, error)
	SetSession(authUser *models.AuthUser) error
}

type sessionService struct {
	repo SessionRepository
}

// New : creates a new core with repository injected
func New(repo SessionRepository) SessionInterface {
	return &sessionService{repo}
}

func (s *sessionService) GetSession(userID string) (*models.Session, error) {
	session, err := s.repo.GetSessionByUserID(userID)
	if err != nil {
		gologger.ERROR("GetSession: error getting the session")
		return nil, mgerror.NewError("error getting the data from database")
	}
	return session, nil
}

func (s *sessionService) SetSession(authUser *models.AuthUser) error {
	session := mapSessionFromAuthUser(authUser)
	err := s.repo.SaveSession(*session)
	if err != nil {
		gologger.ERROR("SetSession: error setting the session on DB")
		return mgerror.NewError("error setting the session on the database")
	}
	return nil
}
