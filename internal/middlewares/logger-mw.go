package middlewares

import (
	"os"

	"github.com/jamolpe/gologger/pkg/models"
)

func logLevels() models.DisplayConfiguration {
	logLevel := os.Getenv("SESSION_COLLECTION")
	result := models.DisplayConfiguration{}
	switch logLevel {
	case "DEV":
		result = models.DisplayConfiguration{
			DisplayDebug:    true,
			DisplayWarnings: true,
			DisplayError:    true,
			DisplayInfo:     true,
		}
	case "PROD":
		result = models.DisplayConfiguration{
			DisplayDebug:    false,
			DisplayWarnings: false,
			DisplayError:    true,
			DisplayInfo:     true,
		}
	default:
		result = models.DisplayConfiguration{
			DisplayDebug:    false,
			DisplayWarnings: true,
			DisplayError:    true,
			DisplayInfo:     true,
		}
	}
	return result
}

func configureLogger(repo models.Repository) models.Configuration {
	config := models.Configuration{
		DisplayLogs: true,
		SaveLogs:    false,
		LogLevels:   logLevels(),
		Repository:  repo,
	}
	return config
}
