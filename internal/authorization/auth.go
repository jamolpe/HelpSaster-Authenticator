package auth

import (
	mgerror "authorization-service/pkg/errors"
	"authorization-service/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func createUserToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = email
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	t, err := token.SignedString([]byte("mytopSecret"))
	if err != nil {
		return "", mgerror.NewError("authorization error")
	}
	return t, nil
}

// Authorization : check if the user is authoriced
func Authorization(authUser *models.User, requestedUser *models.User) (*models.AuthUser, error) {
	logerUser := &models.AuthUser{}
	if authUser.Password == requestedUser.Password {
		token, err := createUserToken(authUser.Email)
		logerUser.User = authUser
		logerUser.Logged = true
		logerUser.Token = token
		return logerUser, err
	}
	return logerUser, nil
}
