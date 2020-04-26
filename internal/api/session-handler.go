package api

import (
	"go-sessioner/pkg/models"

	"github.com/labstack/echo"
)

// CheckValidSession : we check if the session is valid
func (api *API) CheckValidSession(c echo.Context) error {
	cookie, err := c.Cookie("help_saster")
	if err != nil {
		return c.JSON(500, models.ErrorResponse{Code: 005, Message: "cookie not found need to relogin"})
	}
	authUser, err := decodeCookie(cookie)
	if err != nil {
		return c.JSON(422, models.ErrorResponse{Code: 007, Message: "session is not valid"})
	}
	result := api.sessionSrv.CheckValidSession(*authUser)
	return c.JSON(200, result)
}
