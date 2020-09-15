package model

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
	"time"
)

// if upsert company-ru fails, try to company-ru-2 (up to 10 times)
func (c *Company) upsertWithRetry(ctx context.Context) error {
	for i := 1; i <= 10; i += 1 {
		opts := options.Update()
		opts.SetUpsert(true)

		c.UpdatedAt = time.Now().UTC()
		_, err := mongo.Companies.UpdateOne(ctx, bson.M{
			"u": c.URL,
		}, bson.M{
			"$set": c,
		}, opts)
		if err == nil {
			break
		}

		if i == 10 {
			logger.Log.Error().Err(err).Send()
			return err
		}

		c.Slug = strings.Join([]string{
			c.Slug,
			strconv.Itoa(i + 1),
		}, "-")
	}

	return nil
}
