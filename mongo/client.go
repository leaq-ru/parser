package mongo

import (
	"context"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

var (
	Companies  *mongo.Collection
	FileOffset *mongo.Collection
	Cities     *mongo.Collection
)

func init() {
	const timeout = 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().
		SetWriteConcern(writeconcern.New(
			writeconcern.W(1),
			writeconcern.J(true),
		)).
		SetReadConcern(readconcern.Available()).
		SetReadPreference(readpref.SecondaryPreferred()).
		ApplyURI(config.Env.Mongo.URI))
	logger.Must(err)

	err = client.Ping(ctx, nil)
	logger.Must(err)

	parser := client.Database("parser")
	createIndex(parser)

	Companies = parser.Collection("companies")
	FileOffset = parser.Collection("file_offset")
	Cities = parser.Collection("cities")
}
