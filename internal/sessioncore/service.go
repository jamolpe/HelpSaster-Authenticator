package sessioncore

import (
	auth "authorization-service/internal/authorization"
	"authorization-service/pkg/models"

	mgerror "authorization-service/pkg/errors"

	gologger "github.com/jamolpe/go-logger"
)

// SessionInterface : session core service
type SessionInterface interface {
	GetSession(userID string) (*models.Session, error)
	SetSession(authUser *models.AuthUser) error
	CheckValidSession(authUser *models.AuthUser) (bool, error)
}

type sessionService struct {
	repo SessionRepository
}

// New : creates a new core with repository injected
func New(repo SessionRepository) SessionInterface {
	return &sessionService{repo}
}

// GetSession - get and returns the session
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

// CheckValidSession : checks if the session is valid if so we refresh the token
func (s *sessionService) CheckValidSession(authUser *models.AuthUser) (bool, error) {
	session := mapSessionFromAuthUser(authUser)
	session, _ = s.GetSession(session.UserID)
	if session == nil {
		return false, nil
	}
	validation := auth.CheckTokenIsValid(authUser.Token)
	if !validation.IsValid {
		gologger.INFO("CheckSession: token not valid")
		return false, nil
	}
	if validation.IsValid {
		return true, nil
	}
	return true, nil
}
