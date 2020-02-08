package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Router : define the base for the API
func (api *API) Router() {
	e := echo.New()
	defineConfiguration(e)
	api.defineAuthRouter(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func defineConfiguration(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano} method=${method}, uri=${uri}, status=${status} \n",
	}))
}

func (api *API) defineAuthRouter(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("/register", api.Register)
	userGroup.POST("/authenticate", api.Authenticate)
}
