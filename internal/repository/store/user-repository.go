package store

import (
	"go-sessioner/pkg/models"
	"context"
	"errors"
	"os"
	"strings"

	// gologger "github.com/jamolpe/// gologger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createUserCollection(database *mongo.Database) *mongo.Collection {
	userCollectionName := os.Getenv("USERS_COLLECTION")
	userCollection := database.Collection(userCollectionName)
	indexes := []mongo.IndexModel{
		mongo.IndexModel{
			Keys:    bson.D{primitive.E{Key: "email", Value: ""}},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := userCollection.Indexes().CreateMany(context.TODO(), indexes)
	if err != nil && !(strings.Contains(err.Error(), "IndexOptionsConflict")) {
		panic(err.Error())
	}
	return userCollection
}

func (r repository) SaveUser(user models.User) error {
	_, err := r.userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		// gologger.ERROR("Repository: an error ocurred saving new user on the db")
		return errors.New("error creating new user")
	}
	// gologger.INFO("Repository: new user inserted")
	return nil
}

func (r repository) GetUserByEmail(user models.User) (models.User, error) {
	var dbUser models.User
	filter := bson.D{primitive.E{Key: "email", Value: user.Email}}
	err := r.userCollection.FindOne(context.TODO(), filter).Decode(&dbUser)
	if err != nil {
		// gologger.ERROR("Repository: an error ocurred getting the user " + err.Error())
		return dbUser, err
	}
	return dbUser, nil
}
