package main

import (
	"fmt"
	"go-sessioner/internal/api"
	authorizationcore "go-sessioner/internal/authorization-core"
	"go-sessioner/internal/middlewares"
	"go-sessioner/internal/repository/database"
	"go-sessioner/internal/repository/store"
	sessioncore "go-sessioner/internal/session-core"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		panic("error loading environment")
	}
}

func main() {
	client := database.ConfigureAndConnect()
	repo := store.New(client)
	middlewares := middlewares.ConfigureMiddlewares(repo)
	authsrv := authorizationcore.New(repo, middlewares.Logger)
	sessionsrv := sessioncore.New(repo, middlewares.Logger)
	handler := api.New(authsrv, sessionsrv)
	handler.Router()
}
