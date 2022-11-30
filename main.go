package main

import (
	"circleci/apiv2"
	"circleci/config"
	"circleci/db"
	"circleci/fetcher"

	"github.com/sirupsen/logrus"
)

func main() {
	// Config from envars
	config := config.Parse()

	// CircleCI
	circleAPI := apiv2.New(config.APIToken, config.OwnerSlug, config.DefaultTimeout)

	// Database
	mongoDatabase, err := db.New(config.MongoURI, config.DBName)
	if err != nil {
		logrus.Error("unable to connect MongoDB", err)
		return
	}

	// Fetchers
	fetchers := fetcher.New(circleAPI, mongoDatabase, config.DefaultTimeout)
	fetchers.Context()

}
