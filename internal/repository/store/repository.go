package store

import (
	authorizationcore "go-sessioner/internal/authorization-core"
	sessioncore "go-sessioner/internal/session-core"

	"github.com/jamolpe/gologger/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Repository definition
	Repository interface {
		authorizationcore.UserRepository
		sessioncore.SessionRepository
		models.Repository
	}
)

// New : create the new repository
func New(client *mongo.Client) Repository {
	database := client.Database("Clients")
	userCollection := createUserCollection(database)
	sessionCollection := createSessionCollection(database)
	logCollection := createLogCollection(database)
	return &repository{client, database, userCollection, sessionCollection, logCollection}
}

type repository struct {
	client            *mongo.Client
	database          *mongo.Database
	userCollection    *mongo.Collection
	sessionCollection *mongo.Collection
	logCollection     *mongo.Collection
}
