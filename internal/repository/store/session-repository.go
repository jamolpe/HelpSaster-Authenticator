package store

import (
	"authorization-service/pkg/models"
	"context"
	"errors"
	"strings"

	gologger "github.com/jamolpe/go-logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createSessionCollection(database *mongo.Database) *mongo.Collection {
	sessionCollection := database.Collection("Session")
	indexes := []mongo.IndexModel{
		mongo.IndexModel{
			Keys:    bson.D{primitive.E{Key: "createdat", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(1800),
		},
		mongo.IndexModel{
			Keys:    bson.D{primitive.E{Key: "email", Value: ""}},
			Options: options.Index().SetUnique(true),
		},
	}
	// _, err := sessionCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
	// 	Keys:    bson.D{{"createdat", 1}},
	// 	Options: options.Index().SetExpireAfterSeconds(1800),
	// },
	// )
	_, err := sessionCollection.Indexes().CreateMany(context.TODO(), indexes)
	if err != nil && !(strings.Contains(err.Error(), "IndexOptionsConflict")) {
		panic(err.Error())
	}
	return sessionCollection
}

func (r *repository) UpdateSession(session models.Session) error {
	// _, err := r.sessionCollection.UpdateOne(context.TODO(), session)
	_, err := r.sessionCollection.ReplaceOne(context.TODO(), session, session)
	if err != nil {
		gologger.ERROR("Repository: an error ocurred updating the session on the db")
		return errors.New("error creating new session")
	}
	gologger.INFO("Repository: new session inserted")
	return nil
}

func (r *repository) SaveSession(session models.Session) error {
	_, err := r.sessionCollection.InsertOne(context.TODO(), session)
	if err != nil {
		gologger.ERROR("Repository: an error ocurred saving the session on the db")
		return errors.New("error creating new session")
	}
	gologger.INFO("Repository: new session inserted")
	return nil
}

func (r *repository) GetSessionByUserID(UserID string) (*models.Session, error) {
	var dbsession = new(models.Session)
	filter := bson.D{primitive.E{Key: "userid", Value: UserID}}
	err := r.sessionCollection.FindOne(context.TODO(), filter).Decode(&dbsession)
	if err != nil {
		gologger.ERROR("Repository: an error ocurred getting the user " + err.Error())
		return dbsession, err
	}
	return dbsession, nil
}
