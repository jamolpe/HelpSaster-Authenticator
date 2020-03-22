package api

import (
	"authorization-service/pkg/models"
	"fmt"

	"github.com/labstack/echo"
)

// CheckValidSession : we check if the session is valid
func (api *API) CheckValidSession(c echo.Context) error {
	cookie, err := c.Cookie("help_saster")
	if err != nil {
		return c.JSON(500, models.ErrorResponse{Code: 005, Message: "cookie not found need to relogin"})
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.JSON(500, models.ErrorResponse{Code: 006, Message: "session could not be checked due api error"})
}
