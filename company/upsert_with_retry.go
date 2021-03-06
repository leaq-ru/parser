package company

import (
	"context"
	"errors"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
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

	opts := options.Update()
	opts.SetUpsert(true)

	c.WithHash()

	// already have duplicate company with another url
	err = mongo.Companies.FindOne(ctx, bson.M{
		"has": c.Hash,
		"u": bson.M{
			"$ne": c.URL,
		},
	}).Err()
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = nil
		} else {
			logger.Log.Error().Err(err).Send()
			return err
		}
	} else {
		err = errors.New("company with same hash and another url already exists. skip update/insert")
		logger.Log.Error().
			Str("hash", c.Hash).
			Str("url", c.URL).
			Err(err).
			Send()
		return err
	}

	for i := 1; i <= 10; i += 1 {
		c.UpdatedAt = time.Now().UTC()
		_, err := mongo.Companies.UpdateOne(ctx, Company{
			URL: c.URL,
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
