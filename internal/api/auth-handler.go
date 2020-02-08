package api

import (
	"github.com/labstack/echo"
)

// Register : register endpoint
func (api *API) Register(c echo.Context) error {
	return c.String(500, "not implemented")
}

// Authenticate : authenticate endpoint
func (api *API) Authenticate(c echo.Context) error {
	return c.String(500, "not implemented")
}
