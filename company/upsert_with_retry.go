package company

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
	err := c.validate()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return err
	}

	pairs := [][]string{{
		"e", "he",
	}, {
		"p", "hp",
	}, {
		"so.v.g", "hv",
	}, {
		"so.t.u", "ht",
	}, {
		"so.y.u", "hy",
	}, {
		"so.f.u", "hf",
	}, {
		"so.i.u", "hi",
	}, {
		"ap.a.u", "ha",
	}, {
		"ap.g.u", "hg",
	}, {
		"i", "hin",
	}, {
		"k", "hk",
	}, {
		"og", "ho",
	}}

	aggrSetVals := bson.M{}
	for _, pair := range pairs {
		aggrSetVals[pair[1]] = bson.M{
			"$ne": bson.A{"$" + pair[0], nil},
		}
	}
	aggrSetVals["h"] = bson.M{
		"$eq": bson.A{"$h", true},
	}

	opts := options.Update()
	opts.SetUpsert(true)

	for i := 1; i <= 10; i += 1 {
		c.UpdatedAt = time.Now().UTC()
		_, err := mongo.Companies.UpdateOne(ctx, Company{
			URL: c.URL,
		}, bson.A{bson.M{
			"$set": c,
		}, bson.M{
			"$set": aggrSetVals,
		}}, opts)
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
