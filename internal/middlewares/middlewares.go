package middlewares

import (
	"github.com/jamolpe/gologger"
	"github.com/jamolpe/gologger/pkg/models"
)

type Middlewares struct {
	Logger gologger.LoggerI
}

func ConfigureMiddlewares(repo models.Repository) *Middlewares {
	config := configureLogger(repo)
	loggerLibrary := gologger.New(config)
	return &Middlewares{Logger: loggerLibrary}
}
