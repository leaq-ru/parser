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
	DB        *mongo.Database
	Companies *mongo.Collection
	Cities    *mongo.Collection
)

const (
	companies = "companies"
	cities    = "cities"
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
		ApplyURI(config.Env.Mongo.URL))
	logger.Must(err)

	err = client.Ping(ctx, nil)
	logger.Must(err)

	parser := client.Database(config.ServiceName)
	createIndex(parser)

	DB = parser
	Companies = parser.Collection(companies)
	Cities = parser.Collection(cities)
}
