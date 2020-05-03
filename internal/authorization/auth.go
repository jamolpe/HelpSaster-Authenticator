package auth

import (
	mgerror "go-sessioner/pkg/errors"
	"os"
	"strconv"
	"time"

	// gologger "github.com/jamolpe/// gologger"
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"go-sessioner/pkg/models"

	"github.com/dgrijalva/jwt-go"
)

type (
	// ValidationResult validation specification result
	ValidationResult struct {
		IsValid bool
		Expired bool
		Error   bool
	}
	Claims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}
)

var mySigningKey = []byte("mytopSecret")

func GetEmailFromToken(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", mgerror.NewError("Token not valid")
	}
	return claims.Email, nil
}

func createUserToken(email string) (string, error) {
	createdAtInt := int64(time.Second)
	var err error
	var tokenTime int64
	tokenTime, err = strconv.ParseInt(os.Getenv("SESSION_EXPIRATION_TIME"), 10, 64)
	expiresAt := time.Now().Add(time.Duration(tokenTime) * time.Second)
	// Create the Claims
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Subject:   email,
			Issuer:    "gosessioner",
			IssuedAt:  createdAtInt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var ss string
	ss, err = token.SignedString(mySigningKey)
	if err != nil {
		return "", mgerror.NewError("authorization error")
	}
	return ss, nil
}

// CheckTokenIsValid checks if the token is valid
func CheckTokenIsValid(tokenString string) ValidationResult {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if token.Valid {
		return ValidationResult{IsValid: true, Expired: false, Error: false}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return ValidationResult{IsValid: false, Expired: false, Error: true}
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// gologger.ERROR("TokenValidation: token expired")
			return ValidationResult{IsValid: false, Expired: true, Error: false}
		} else {
			// gologger.ERROR("TokenValidation: couldn't handle this token:" + err.Error())
			return ValidationResult{IsValid: false, Expired: false, Error: true}
		}
	}
	return ValidationResult{IsValid: false, Expired: false, Error: true}
}

// SecureString : Secure the password using hash
func SecureString(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckCorrespondingString : check if the plain password correspond with the hash
func CheckCorrespondingString(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Authorization : check if the user is authoriced
func Authorization(dbUser *models.User, requestedUser *models.User) (*models.AuthUser, string, error) {
	logerUser := &models.AuthUser{}
	if CheckCorrespondingString(requestedUser.Password, dbUser.Password) {
		token, err := createUserToken(dbUser.Email)
		logerUser.User = dbUser
		logerUser.Logged = true
		return logerUser, token, err
	}
	return logerUser, "", nil
}
