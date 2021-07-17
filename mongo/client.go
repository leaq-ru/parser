package mongo

import (
	"context"
	"time"

	"github.com/leaq-ru/parser/config"
	"github.com/leaq-ru/parser/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var (
	Client               *mongo.Client
	Companies            *mongo.Collection
	Posts                *mongo.Collection
	CachedLists          *mongo.Collection
	Reviews              *mongo.Collection
	Categories           *mongo.Collection
	Cities               *mongo.Collection
	Technologies         *mongo.Collection
	TechnologyCategories *mongo.Collection
	DNS                  *mongo.Collection
)

const (
	companies            = "companies"
	posts                = "posts"
	cachedLists          = "cached_lists"
	reviews              = "reviews"
	categories           = "categories"
	cities               = "cities"
	technologies         = "technologies"
	technologyCategories = "technology_categories"
	dns                  = "dns"
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
	Reviews = parser.Collection(reviews)
	Categories = parser.Collection(categories)
	Cities = parser.Collection(cities)
	Technologies = parser.Collection(technologies)
	TechnologyCategories = parser.Collection(technologyCategories)
	DNS = parser.Collection(dns)
}
