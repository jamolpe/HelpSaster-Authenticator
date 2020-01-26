package core

import (
	"errors"

	gologger "github.com/jamolpe/go-logger"
)

// UserServiceInterface : authentication core service
type UserServiceInterface interface {
	UserRegister(user *User) (bool, error)
	Authenticate(loginUserData LoginUserData) (*LoginUserResponse, error)
}

type userSrv struct {
	repo UserRepository
}

// NewUserService : creates a new core with repository injected
func NewUserService(repo UserRepository) UserServiceInterface {
	return &userSrv{repo}
}

func (s *userSrv) UserRegister(user *User) (bool, error) {
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

func (s *userSrv) Authenticate(LoginUserData LoginUserData) (*LoginUserResponse, error) {
	return &LoginUserResponse{}, nil
}
