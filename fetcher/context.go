package fetcher

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (fetcher Fetcher) Context() {
	log := logrus.WithField("fetch", "context")

	// get from api
	contextResp, err := fetcher.circleAPI.ListContext("")
	circleAPILog := log.WithField("operation", "apiv2.ListContext")
	if err != nil {
		circleAPILog.Error(err)
		return
	}
	circleAPILog.Info(contextResp)

	if len(contextResp.Items) == 0 {
		log.Info("no items to insert")
		return
	}

	// insert to database
	contextRepo := fetcher.database.Collection("context")
	ctx, cancel := context.WithTimeout(context.Background(), fetcher.operationTimeout)
	defer cancel()
	for _, item := range contextResp.Items {
		contextRepoLog := log.
			WithField("operation", "contextRepo.Upsert").
			WithField("_id", item.ID)

		_, err := contextRepo.UpdateOne(ctx, bson.M{
			"_id": item.ID,
		}, bson.M{
			"$set": item,
		}, options.Update().SetUpsert(true))

		if err != nil {
			contextRepoLog.Error(err)
			continue
		}

		contextRepoLog.Info("success")
	}

}
