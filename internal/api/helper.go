package api

import (
	"go-sessioner/pkg/models"
)

func mapRegisterUserToModel(userToRegister *models.RegisterUser) *models.User {
	var user = new(models.User)
	user.Email = userToRegister.Email
	user.Name = userToRegister.Name
	user.Password = userToRegister.Password
	user.Age = userToRegister.Age
	return user
}

func mapLoginUserToModel(userToLogin *models.LoginUser) *models.User {
	var user = new(models.User)
	user.Email = userToLogin.Email
	user.Password = userToLogin.Password
	return user
}
