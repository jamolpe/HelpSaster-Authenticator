package sessioncore

import (
	auth "go-sessioner/internal/authorization"
	"go-sessioner/pkg/models"

	mgerror "go-sessioner/pkg/errors"

	"github.com/jamolpe/gologger"
)

// SessionInterface : session core service
type SessionInterface interface {
	GetSession(userID string) (*models.Session, error)
	SetSession(authUser *models.AuthUser) error
	CheckValidSession(authUser *models.AuthUser) (bool, error)
}

type sessionService struct {
	repo   SessionRepository
	logger gologger.LoggerI
}

// New : creates a new core with repository injected
func New(repo SessionRepository, logger gologger.LoggerI) SessionInterface {
	return &sessionService{repo, logger}
}

// GetSession - get and returns the session
func (s *sessionService) GetSession(userID string) (*models.Session, error) {
	session, err := s.repo.GetSessionByUserID(userID)
	if err != nil {
		s.logger.ERROR("GetSession: error getting the session")
		return nil, mgerror.NewError("error getting the data from database")
	}
	return session, nil
}

func (s *sessionService) modifyWithExistingSession(session *models.Session) {
	existingSession, _ := s.GetSession(session.UserID)
	if existingSession != nil {
		session = existingSession
	}
}

func (s *sessionService) SetSession(authUser *models.AuthUser) error {
	session := mapSessionFromAuthUser(authUser)
	s.modifyWithExistingSession(session)
	err := s.repo.UpdateSession(*session)
	if err != nil {
		s.logger.ERROR("SetSession: error setting the session on DB")
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
		s.logger.INFO("CheckSession: token not valid")
		return false, nil
	}
	if validation.IsValid {
		return true, nil
	}
	return true, nil
}
