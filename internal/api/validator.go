package api

import (
	"authorization-service/pkg/models"
	"regexp"
	"unicode"
)

var re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func validateRegisterUser(userToRegister *models.RegisterUser) (bool, string) {
	if userToRegister.Age > 150 || userToRegister.Age < 0 {
		return false, "Age is not valid"
	}
	if !re.MatchString(userToRegister.Email) {
		return false, "Email is not valid"
	}

	if userToRegister.LastName == "" {
		return false, "LastName can not be empty"
	}

	if userToRegister.Name == "" {
		return false, "Name can not be empty"
	}

	if !validatePassword(userToRegister.Password) {
		return false, "Password is not valid(upper, lower, number and special case)"
	}

	return true, ""
}

func validatePassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func validateAuthenticateUser(loginUser *models.LoginUser) (bool, string) {
	if !re.MatchString(loginUser.Email) {
		return false, "email is not valid"
	}
	if !validatePassword(loginUser.Password) {
		return false, "password is not valid"
	}
	return true, ""
}
