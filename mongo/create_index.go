package mongo

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndex(db *mongo.Database) {
	compURLIndex := options.Index()
	compURLIndex.SetUnique(true)
	compSlugIndex := options.Index()
	compSlugIndex.SetUnique(true)
	companies := db.Collection(companies)
	_, err := companies.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{
				"u": 1,
			},
			Options: compURLIndex,
		},
		{
			Keys: bson.M{
				"s": 1,
			},
			Options: compSlugIndex,
		},
		{
			Keys: bson.M{
				"l.c": 1,
			},
		},
		{
			Keys: bson.M{
				"c": 1,
			},
		},
	})
	logger.Must(err)
}
