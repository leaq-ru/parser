package company

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/ptr"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func companySetOffline(ctx context.Context, url string) (err error) {
	_, err = mongo.Companies.UpdateOne(ctx, Company{
		URL: url,
	}, bson.M{
		"$set": Company{
			Online:    ptr.Bool(false),
			UpdatedAt: time.Now().UTC(),
		},
	})
	logger.Err(err)
	return err
}
