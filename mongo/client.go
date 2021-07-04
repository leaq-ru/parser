package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	companies   = "companies"
	posts       = "posts"
	cachedLists = "cached_lists"
)

type StartSession = func(...*options.SessionOptions) (mongo.Session, error)

func NewClient(
	dbName string,
	url string,
) (
	StartSession,
	*mongo.Collection,
	*mongo.Collection,
	*mongo.Collection,
	error,
) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().
		SetWriteConcern(writeconcern.New(
			writeconcern.WMajority(),
			writeconcern.J(true),
		)).
		SetReadConcern(readconcern.Majority()).
		SetReadPreference(readpref.SecondaryPreferred()).
		ApplyURI(url))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	db := client.Database(dbName)

	err = createIndex(db)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return client.StartSession,
		db.Collection(companies),
		db.Collection(posts),
		db.Collection(cachedLists),
		nil
}
