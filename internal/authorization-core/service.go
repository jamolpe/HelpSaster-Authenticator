package authorizationcore

import (
	"errors"
	auth "go-sessioner/internal/authorization"
	"go-sessioner/pkg/models"
)

// AuthServiceInterface : authentication core service
type AuthServiceInterface interface {
	UserRegister(user *models.User) (bool, error)
	Authenticate(loginUserData *models.User) (bool, *models.AuthUser, error)
}

type authService struct {
	repo UserRepository
}

// New : creates a new core with repository injected
func New(repo UserRepository) AuthServiceInterface {
	return &authService{repo}
}

// UserRegister : register the user in the databa
func (s *authService) UserRegister(user *models.User) (bool, error) {
	exist := s.findIfUserExist(user)
	if exist {
		//// gologger.INFO("Register: user already exist")
		return false, errors.New("user already exist")
	}
	user.Password, _ = auth.SecureString(user.Password)
	err := s.repo.SaveUser(*user)
	if err != nil {
		//// gologger.ERROR(`save new user error`)
		return false, err
	}
	return true, nil
}

// Authenticate : authenticate the user and give a token
func (s *authService) Authenticate(loginUser *models.User) (bool, *models.AuthUser, error) {
	user := &models.User{Email: loginUser.Email}
	dbUser := s.getUserFromDatabase(*user)
	if *dbUser == (models.User{}) {
		//// gologger.INFO("Authenticate: user not found")
		return false, nil, nil
	}
	sessionUser, err := auth.Authorization(dbUser, loginUser)
	if err != nil {
		//// gologger.ERROR("Authenticate: " + err.Error())
		return false, nil, err
	}
	if *sessionUser == (models.AuthUser{}) {
		//// gologger.DEBUG("Authenticate: wrong password")
		return false, sessionUser, nil
	}
	//// gologger.INFO("Authenticate: user authentified " + loginUser.Email)
	return true, sessionUser, nil
}
