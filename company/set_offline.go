package company

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func companySetOffline(ctx context.Context, slug string) (err error) {
	_, err = mongo.Companies.UpdateOne(ctx, bson.M{
		"s": slug,
	}, bson.M{
		"$set": bson.M{
			"ua": time.Now().UTC(),
		},
		"$unset": bson.M{
			"o": "",
		},
	})
	logger.Err(err)
	return err
}
