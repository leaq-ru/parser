package post

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Get(ctx context.Context, companyID primitive.ObjectID, skip, limit uint32, excludeIDs []primitive.ObjectID) (
	res []Post,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := bson.M{
		"c": companyID,
	}
	if len(excludeIDs) != 0 {
		query["_id"] = bson.M{
			"$nin": excludeIDs,
		}
	}

	cur, err := mongo.Posts.Find(ctx, query, options.Find().
		SetSort(bson.M{
			"d": -1,
		}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit)))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	err = cur.All(ctx, &res)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return
}
