package sessioncore

import (
	auth "go-sessioner/internal/authorization"
	"go-sessioner/pkg/models"

	mgerror "go-sessioner/pkg/errors"

	guuid "github.com/google/uuid"

	"github.com/jamolpe/gologger"
)

// SessionInterface : session core service
type SessionInterface interface {
	GetSession(ID guuid.UUID) (*models.Session, error)
	SetSession(session models.Session) error
	CheckValidSession(session models.Session) bool
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
func (s *sessionService) GetSession(sessionID guuid.UUID) (*models.Session, error) {
	session, err := s.repo.GetSessionByID(sessionID.String())
	if err != nil {
		s.logger.ERROR("GetSession: error getting the session")
		return nil, mgerror.NewError("error getting the data from database")
	}
	return session, nil
}

func (s *sessionService) modifyWithExistingSession(session *models.Session) bool {
	existingSession, _ := s.GetSession(session.ID)
	if existingSession != nil {
		session = existingSession
		return true
	}
	return false
}

func (s *sessionService) SetSession(session models.Session) error {
	modified := s.modifyWithExistingSession(&session)
	var err error
	if modified {
		err = s.repo.UpdateSession(session)
	} else {
		err = s.repo.SaveSession(session)
	}
	if err != nil {
		s.logger.ERROR("SetSession: error setting the session on DB")
		return mgerror.NewError("error setting the session on the database")
	}
	return nil
}

// CheckValidSession : checks if the session is valid if so we refresh the token
func (s *sessionService) CheckValidSession(session models.Session) bool {
	validation := auth.CheckTokenIsValid(session.Token)
	if !validation.IsValid {
		s.logger.INFO("CheckSession: token not valid")
		return false
	}
	if validation.IsValid {
		return true
	}
	return false
}
