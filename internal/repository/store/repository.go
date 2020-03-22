package store

import (
	authorizationcore "authorization-service/internal/authorization-core"
	sessioncore "authorization-service/internal/session-core"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Repository definition
	Repository interface {
		authorizationcore.UserRepository
		sessioncore.SessionRepository
	}
)

// New : create the new repository
func New(client *mongo.Client) Repository {
	database := client.Database("Clients")
	userCollection := createUserCollection(database)
	sessionCollection := createSessionCollection(database)
	return &repository{client, database, userCollection, sessionCollection}
}

type repository struct {
	client            *mongo.Client
	database          *mongo.Database
	userCollection    *mongo.Collection
	sessionCollection *mongo.Collection
}
