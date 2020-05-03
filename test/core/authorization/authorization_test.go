package authorizationtest

import (
	auth "go-sessioner/internal/authorization"
	"go-sessioner/pkg/models"
	"testing"
)

func compareSessionUser(user1 *models.AuthUser, user2 *models.AuthUser) bool {
	var result = true
	if user1.Logged != user2.Logged {
		result = false
	}

	// should be the same pointer
	if user1.User != user2.User {
		result = false
	}

	return result
}
func Test_Authorization_WrongPassword(t *testing.T) {
	user := &models.User{Email: "email@email.com", Password: "diferentPassword"}
	loginUser := &models.User{Email: "email@email.com", Password: "anotherPassword"}

	result, _, _ := auth.Authorization(user, loginUser)
	if *result != (models.AuthUser{}) {
		t.Error("password are equals")
	}
}

func Test_Authorization_CorrectPassword(t *testing.T) {
	user := &models.User{Email: "email@email.com", Password: "samePassword"}
	loginUser := &models.User{Email: "email@email.com", Password: "samePassword"}

	expectedResult := &models.AuthUser{User: user, Logged: true}
	user.Password, _ = auth.SecureString(user.Password)
	result, _, _ := auth.Authorization(user, loginUser)
	if ok := compareSessionUser(expectedResult, result); !ok {
		t.Error("wrong session user")
	}
}
