package category

import (
	"context"
	"errors"
	"github.com/gosimple/slug"
	"github.com/jbrukh/bayesian"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
)

func (c *Category) GetOrCreate(ctx context.Context, title bayesian.Class) (doc Category, err error) {
	doc = Category{
		Title: title,
		Slug:  slug.Make(string(title)),
	}

	id, err := mongo.Categories.InsertOne(ctx, doc)
	if err != nil {
		e := mongo.Categories.FindOne(ctx, bson.M{
			"s": doc.Slug,
		}).Decode(&doc)
		if e != nil {
			if errors.Is(e, m.ErrNoDocuments) {
				logger.Log.Error().Err(e).Err(err).Msg("insert error and no docs found")
				return
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
