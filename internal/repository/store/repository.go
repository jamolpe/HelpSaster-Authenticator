package store

import (
	"authorization-service/internal/core"

	"go.mongodb.org/mongo-driver/mongo"
)

// NewStore : create the new repository
func NewStore(client *mongo.Client) core.Repository {
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
