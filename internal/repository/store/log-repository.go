package store

import (
	"context"
	"os"

	"github.com/jamolpe/gologger/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func createLogCollection(database *mongo.Database) *mongo.Collection {
	logCollectionName := os.Getenv("LOG_COLLECTION")
	logCollection := database.Collection(logCollectionName)
	return logCollection
}

func (r *repository) SaveLog(log models.LogModel) error {
	_, err := r.logCollection.InsertOne(context.TODO(), log)
	return err
}
