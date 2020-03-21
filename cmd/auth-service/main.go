package main

import (
	"authorization-service/internal/api"
	authorizationcore "authorization-service/internal/authorization-core"
	"authorization-service/internal/repository/database"
	"authorization-service/internal/repository/store"
	sessioncore "authorization-service/internal/session-core"
)

func main() {
	client := database.ConfigureAndConnect()
	repo := store.New(client)
	authsrv := authorizationcore.New(repo)
	sessionsrv := sessioncore.New(repo)
	handler := api.New(authsrv, sessionsrv)
	handler.Router()
}
