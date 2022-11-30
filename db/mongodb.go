package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(mongoURI string, dbName string) (*mongo.Database, error) {
	logrus.Info("creating mongoDB client...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Error("unable to connect MongoDB")
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		logrus.Error("unable to ping MongoDB")
		return nil, err
	}

	mongodbDatabase := mongoClient.Database(dbName)
	return mongodbDatabase, nil
}
