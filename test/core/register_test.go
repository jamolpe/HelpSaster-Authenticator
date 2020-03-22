package coretest

import (
	authorizationcore "authorization-service/internal/authorization-core"
	"authorization-service/pkg/models"
	"errors"
	"testing"
)

// NewRepository : return new stub
func NewRepository(resultSaveUser error, resultGetUser models.User, resultGetUserError error) authorizationcore.UserRepository {
	return &repositoryStub{resultSaveUser, resultGetUser, resultGetUserError}
}

type repositoryStub struct {
	resultSaveUser     error
	resultGetUser      models.User
	resultGetUserError error
}

func (r repositoryStub) SaveUser(user models.User) error {
	return r.resultSaveUser
}

func (r repositoryStub) GetUserByEmail(user models.User) (models.User, error) {
	return r.resultGetUser, r.resultGetUserError
}

func Test_RegisterWithError(t *testing.T) {
	repo := NewRepository(errors.New("error creating new user"), models.User{}, nil)
	srv := authorizationcore.New(repo)
	expectedResult := false
	expectedError := errors.New("error creating new user")
	result, err := srv.UserRegister(&models.User{})
	if result != expectedResult {
		t.Error("saved user not posible")
	}
	if err.Error() != expectedError.Error() {
		t.Error("saved user no error ocurred")
	}
}

func Test_RegisterUserExistingUser(t *testing.T) {
	repo := NewRepository(nil, models.User{Email: "email@gmail.com", Password: "123", Name: "name", Age: 18}, nil)
	srv := authorizationcore.New(repo)
	expectedResult := false
	expectedError := errors.New("user already exist")
	result, err := srv.UserRegister(&models.User{Email: "email@gmail.com", Password: "123", Name: "name", Age: 18})
	if result != expectedResult {
		t.Error("saved user not posible")
	}
	if err.Error() != expectedError.Error() {
		t.Error("saved user no error ocurred")
	}
}

func Test_RegisterUserNoUsers(t *testing.T) {
	repo := NewRepository(nil, models.User{}, nil)
	srv := authorizationcore.New(repo)
	expectedResult := true
	result, err := srv.UserRegister(&models.User{Email: "email@gmail.com", Password: "123", Name: "name", Age: 18})
	if result != expectedResult {
		t.Error("saved user not posible")
	}
	if err != nil {
		t.Error("saved user error ocurred")
	}
}
