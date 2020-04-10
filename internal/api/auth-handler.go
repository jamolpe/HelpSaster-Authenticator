package api

import (
	"go-sessioner/pkg/models"
	"os"

	"github.com/labstack/echo"
)

// Register : register endpoint
func (api *API) Register(c echo.Context) error {
	userToRegister := new(models.RegisterUser)
	if err := c.Bind(userToRegister); err == nil {
		if validation, message := validateRegisterUser(userToRegister); !validation {
			return c.JSON(422, models.ErrorResponse{Code: 001, Message: message})
		}
		user := mapRegisterUserToModel(userToRegister)
		registered, err := api.authSrv.UserRegister(user)
		if err == nil && registered {
			return c.JSON(200, models.RegisterResponse{Message: "user registered"})
		}
		return c.JSON(422, models.ErrorResponse{Code: 002, Message: "user not registered due an error " + err.Error()})
	}
	return c.JSON(500, models.ErrorResponse{Code: 003, Message: "user not registered due an api error"})
}

// Authenticate : authenticate endpoint
func (api *API) Authenticate(c echo.Context) error {
	userToLogin := new(models.LoginUser)
	if err := c.Bind(userToLogin); err == nil {
		if validation, message := validateAuthenticateUser(userToLogin); !validation {
			return c.JSON(500, models.ErrorResponse{Code: 005, Message: message})
		}
		user := mapLoginUserToModel(userToLogin)
		authenticated, authUser, err := api.authSrv.Authenticate(user)
		if err == nil && authenticated == false {
			return c.JSON(422, models.ErrorResponse{Code: 06, Message: "login incorrect"})
		}
		if err == nil && authenticated {
			sessionCookie, cookieerr := createCookie(*authUser)
			if cookieerr != nil {
				return c.JSON(500, models.ErrorResponse{Code: 005, Message: "error creating cookie"})
			}
			if os.Getenv("GO_ENV") != "test" {
				go api.sessionSrv.SetSession(authUser)
			}
			c.SetCookie(sessionCookie)
			return c.JSON(200, models.Authresponse{Message: "user authorized", LogedUser: *authUser})
		}
	}
	return c.JSON(500, models.ErrorResponse{Code: 004, Message: "user not logged due api error"})
}
