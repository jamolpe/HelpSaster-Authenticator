package core

import (
	auth "authorization-service/internal/core/authorization"
	"authorization-service/pkg/models"
	"errors"

	gologger "github.com/jamolpe/go-logger"
)

// AuthServiceInterface : authentication core service
type AuthServiceInterface interface {
	UserRegister(user *models.User) (bool, error)
	Authenticate(loginUserData models.LoginUser) (*models.SessionUser, error)
}

type authSrv struct {
	repo UserRepository
}

// NewUserService : creates a new core with repository injected
func NewUserService(repo UserRepository) AuthServiceInterface {
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
func (s *authSrv) Authenticate(loginUser models.LoginUser) (*models.SessionUser, error) {
	user := &models.User{Email: loginUser.Email}
	dbUser := s.getUserFromDatabase(*user)
	sessionUser, err := auth.Authorization(dbUser, loginUser)
	if err != nil {
		gologger.ERROR("Authenticate: " + err.Error())
		return nil, err
	}
	gologger.INFO("Authenticate: user authentified " + loginUser.Email)
	return sessionUser, nil
}
