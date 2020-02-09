package store

import (
	"authorization-service/internal/core"
	"authorization-service/pkg/models"
	"context"
	"errors"

	gologger "github.com/jamolpe/go-logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewStore : create the new repository
func NewStore(client *mongo.Client) core.UserRepository {
	database := client.Database("Clients")
	userCollection := database.Collection("Users")
	return &repository{client, database, userCollection}
}

type repository struct {
	client         *mongo.Client
	database       *mongo.Database
	userCollection *mongo.Collection
}

func (r repository) SaveUser(user models.User) error {
	_, err := r.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		gologger.ERROR("Repository: an error ocurred saving new user on the db")
		return errors.New("error creating new user")
	}
	gologger.INFO("Repository: new user inserted")
	return nil
}

func (r repository) GetUserByEmail(user models.User) (models.User, error) {
	var dbUser models.User
	filter := bson.D{{"email", user.Email}}
	err := r.userCollection.FindOne(context.TODO(), filter).Decode(&dbUser)
	if err != nil {
		gologger.ERROR("Repository: an error ocurred getting the user " + err.Error())
		return dbUser, err
	}
	return dbUser, nil
}
