package company

import (
	"context"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func companySetOffline(ctx context.Context, url string) (err error) {
	_, err = mongo.Companies.UpdateOne(ctx, Company{
		URL: url,
	}, bson.M{
		"$set": Company{
			UpdatedAt: time.Now().UTC(),
		},
		"$unset": bson.M{
			"o": 1,
		},
	})
	logger.Err(err)
	return err
}
