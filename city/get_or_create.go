package city

import (
	"context"
	"errors"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongod "go.mongodb.org/mongo-driver/mongo"
)

func (c *City) GetOrCreate(ctx context.Context, title NormalCaseCity) (doc City, err error) {
	doc = City{
		Title: title,
		Slug:  slug.Make(string(title)),
	}

	id, err := mongo.Cities.InsertOne(ctx, doc)
	if err != nil {
		e := mongo.Cities.FindOne(ctx, bson.M{
			"s": doc.Slug,
		}).Decode(&doc)
		if e != nil {
			if errors.Is(e, mongod.ErrNoDocuments) {
				logger.Log.Panic().Err(e).Err(err).Msg("insert error and no docs found")
			}
		}
		err = nil
		return
	}

	oID, ok := id.InsertedID.(primitive.ObjectID)
	if !ok {
		msg := "failed cast to primitive.ObjectID"
		logger.Log.Error().Interface("id", id).Msg(msg)
		err = errors.New(msg)
		return
	}
	doc.ID = oID
	return
}
