package api

import (
	"authorization-service/internal/core"
)

// New : define a new api
func New(authSrv core.AuthServiceInterface) *API {
	return &API{
		authSrv: authSrv,
	}
}

type (
	// API : core endpoint
	API struct {
		authSrv core.AuthServiceInterface
	}
)
