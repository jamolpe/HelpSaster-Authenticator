package coretest

import (
	authorizationcore "authorization-service/internal/authorization-core"
	"authorization-service/pkg/models"
	"testing"
)

func Test_Authenticate_NotExistingUser(t *testing.T) {
	loginUser := &models.User{Email: "user@user.com", Password: "password"}
	repo := NewRepository(nil, models.User{}, nil)
	srv := authorizationcore.New(repo)

	result, session, err := srv.Authenticate(loginUser)
	if result != false {
		t.Error("fail user authenticated")
	}
	if session != nil {
		t.Error("fail user session has info")
	}
	if err != nil {
		t.Error("fail user error happened")
	}
}
