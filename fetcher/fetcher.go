package fetcher

import (
	"circleci/apiv2"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Fetcher struct {
	circleAPI        apiv2.API
	database         *mongo.Database
	operationTimeout time.Duration
}

func New(apiV2 apiv2.API, mongoDatabase *mongo.Database, operationTimeout time.Duration) Fetcher {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return Fetcher{
		circleAPI:        apiV2,
		database:         mongoDatabase,
		operationTimeout: operationTimeout,
	}
}
