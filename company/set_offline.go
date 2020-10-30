package company

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func companySetOffline(ctx context.Context, url string) (err error) {
	_, err = mongo.Companies.UpdateOne(ctx, Company{
		URL: url,
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
