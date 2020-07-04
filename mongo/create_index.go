package mongo

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndex(db *mongo.Database) {
	companies := db.Collection("companies")
	optsURL := &options.IndexOptions{}
	optsURL.SetUnique(true)
	_, err := companies.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{
				"u": 1,
			},
			Options: optsURL,
		},
	})
	logger.Must(err)
}
