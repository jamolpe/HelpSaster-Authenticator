package main

import (
	"authorization-service/internal/api"
	"authorization-service/internal/core"
	"authorization-service/internal/repository/database"
	"authorization-service/internal/repository/store"
)

func main() {
	client := database.ConfigureAndConnect()
	repo := store.NewStore(client)
	srv := core.NewUserService(repo)
	handler := api.New(srv)
	handler.Router()
}
