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
	Client      *mongo.Client
	Companies   *mongo.Collection
	Posts       *mongo.Collection
	CachedLists *mongo.Collection
)

const (
	companies   = "companies"
	posts       = "posts"
	cachedLists = "cached_lists"
)

func init() {
	if config.Env.MongoDB.URL == "" {
		return
	}

	const timeout = 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().
		SetWriteConcern(writeconcern.New(
			writeconcern.WMajority(),
			writeconcern.J(true),
		)).
		SetReadConcern(readconcern.Majority()).
		SetReadPreference(readpref.SecondaryPreferred()).
		ApplyURI(config.Env.MongoDB.URL))
	logger.Must(err)

	err = client.Ping(ctx, nil)
	logger.Must(err)

	parser := client.Database(config.ServiceName)
	createIndex(parser)

	Client = parser.Client()
	Companies = parser.Collection(companies)
	Posts = parser.Collection(posts)
	CachedLists = parser.Collection(cachedLists)
}
