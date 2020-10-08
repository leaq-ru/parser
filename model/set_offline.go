package model

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func companySetOffline(ctx context.Context, slug string) (err error) {
	_, err = mongo.Companies.UpdateOne(ctx, bson.M{
		"s": slug,
	}, bson.M{
		"$unset": bson.M{
			"o": "",
		},
	})
	logger.Err(err)
	return err
}