package core

import (
	auth "authorization-service/internal/core/authorization"
	mgerror "authorization-service/pkg/errors"
	"authorization-service/pkg/models"
	"errors"

	gologger "github.com/jamolpe/go-logger"
)

// AuthServiceInterface : authentication core service
type AuthServiceInterface interface {
	UserRegister(user *models.User) (bool, error)
	Authenticate(loginUserData *models.User) (bool, *models.AuthUser, error)
	GetSession(userID string) (*models.Session, error)
	SetSession(authUser *models.AuthUser) error
}

type authSrv struct {
	repo Repository
}

// New : creates a new core with repository injected
func New(repo Repository) AuthServiceInterface {
	return &authSrv{repo}
}

// UserRegister : register the user in the databa
func (s *authSrv) UserRegister(user *models.User) (bool, error) {
	exist := s.findIfUserExist(user)
	if exist {
		gologger.INFO("Register: user already exist")
		return false, errors.New("user already exist")
	}
	err := s.repo.SaveUser(*user)
	if err != nil {
		gologger.ERROR(`save new user error`)
		return false, err
	}
	return true, nil
}

// Authenticate : authenticate the user and give a token
func (s *authSrv) Authenticate(loginUser *models.User) (bool, *models.AuthUser, error) {
	user := &models.User{Email: loginUser.Email}
	dbUser := s.getUserFromDatabase(*user)
	if *dbUser == (models.User{}) {
		gologger.INFO("Authenticate: user not found")
		return false, nil, nil
	}
	sessionUser, err := auth.Authorization(dbUser, loginUser)
	if err != nil {
		gologger.ERROR("Authenticate: " + err.Error())
		return false, nil, err
	}
	if *sessionUser == (models.AuthUser{}) {
		gologger.DEBUG("Authenticate: wrong password")
		return false, sessionUser, nil
	}
	gologger.INFO("Authenticate: user authentified " + loginUser.Email)
	return true, sessionUser, nil
}

func (s *authSrv) GetSession(userID string) (*models.Session, error) {
	session, err := s.repo.GetSessionByUserID(userID)
	if err != nil {
		gologger.ERROR("GetSession: error getting the session")
		return nil, mgerror.NewError("error getting the data from database")
	}
	return session, nil
}

func (s *authSrv) SetSession(authUser *models.AuthUser) error {
	session := mapSessionFromAuthUser(authUser)
	err := s.repo.SaveSession(*session)
	if err != nil {
		gologger.ERROR("SetSession: error setting the session on DB")
		return mgerror.NewError("error setting the session on the database")
	}
	return nil
}
