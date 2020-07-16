package company

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

// if upsert company-ru fails, try to company-ru-2 (up to 10 times)
func (c *Company) upsertWithRetry(ctx context.Context) error {
	opts := options.Update()
	opts.SetUpsert(true)

	for i := 0; i < 10; i += 1 {
		_, err := mongo.Companies.UpdateOne(ctx, bson.M{
			"u": c.URL,
		}, bson.M{
			"$set": c,
		}, opts)

		if err == nil {
			break
		}

		if i == 9 {
			logger.Log.Error().Err(err).Send()
			return err
		}

		c.Slug = strings.Join([]string{
			c.Slug,
			strconv.Itoa(i + 2),
		}, "-")
	}

	return nil
}
