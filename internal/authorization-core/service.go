package authorizationcore

import (
	"errors"
	auth "go-sessioner/internal/authorization"
	"go-sessioner/pkg/models"

	"github.com/jamolpe/gologger"
)

// AuthServiceInterface : authentication core service
type AuthServiceInterface interface {
	UserRegister(user *models.User) (bool, error)
	Authenticate(loginUserData *models.User) (bool, *models.AuthUser, string, error)
}

type authService struct {
	repo   UserRepository
	logger gologger.LoggerI
}

// New : creates a new core with repository injected
func New(repo UserRepository, logger gologger.LoggerI) AuthServiceInterface {
	return &authService{repo, logger}
}

// UserRegister : register the user in the databa
func (s *authService) UserRegister(user *models.User) (bool, error) {
	exist := s.findIfUserExist(user)
	if exist {
		s.logger.INFO("Register: user already exist")
		return false, errors.New("user already exist")
	}
	user.Password, _ = auth.SecureString(user.Password)
	err := s.repo.SaveUser(*user)
	if err != nil {
		s.logger.ERROR(`save new user error`)
		return false, err
	}
	return true, nil
}

// Authenticate : authenticate the user and give a token
func (s *authService) Authenticate(loginUser *models.User) (bool, *models.AuthUser, string, error) {
	user := &models.User{Email: loginUser.Email}
	dbUser := s.getUserFromDatabase(*user)
	if *dbUser == (models.User{}) {
		s.logger.INFO("Authenticate: user not found")
		return false, nil, "", nil
	}
	authoricedUser, token, err := auth.Authorization(dbUser, loginUser)
	if err != nil {
		s.logger.ERROR("Authenticate: " + err.Error())
		return false, nil, "", err
	}
	if *authoricedUser == (models.AuthUser{}) {
		s.logger.DEBUG("Authenticate: wrong password")
		return false, authoricedUser, "", nil
	}
	s.logger.INFO("Authenticate: user authentified " + loginUser.Email)
	return true, authoricedUser, token, nil
}
