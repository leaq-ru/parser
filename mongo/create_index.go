package mongo

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func createIndex(db *m.Database) {
	ctx := context.Background()

	_, err := db.Collection(companies).Indexes().CreateMany(ctx, []m.IndexModel{{
		Keys: bson.M{
			"u": 1,
		},
		Options: options.Index().SetUnique(true),
	}, {
		Keys: bson.M{
			"s": 1,
		},
		Options: options.Index().SetUnique(true),
	}, {
		Keys: bson.D{{
			Key:   "has",
			Value: 1,
		}, {
			Key:   "u",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "l.c",
			Value: 1,
		}, {
			Key:   "c",
			Value: 1,
		}, {
			Key:   "h",
			Value: 1,
		}, {
			Key:   "pd",
			Value: -1,
		}, {
			Key:   "_id",
			Value: -1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "l.c",
			Value: 1,
		}, {
			Key:   "h",
			Value: 1,
		}, {
			Key:   "pd",
			Value: -1,
		}, {
			Key:   "_id",
			Value: -1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "c",
			Value: 1,
		}, {
			Key:   "h",
			Value: 1,
		}, {
			Key:   "pd",
			Value: -1,
		}, {
			Key:   "_id",
			Value: -1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "ti",
			Value: 1,
		}, {
			Key:   "h",
			Value: 1,
		}, {
			Key:   "pd",
			Value: -1,
		}, {
			Key:   "_id",
			Value: -1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "h",
			Value: 1,
		}, {
			Key:   "pd",
			Value: -1,
		}, {
			Key:   "_id",
			Value: -1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "pd",
			Value: 1,
		}},
	}})
	logger.Must(err)

	_, err = db.Collection(posts).Indexes().CreateOne(ctx, m.IndexModel{
		Keys: bson.D{{
			Key:   "c",
			Value: 1,
		}, {
			Key:   "d",
			Value: -1,
		}},
	})
	logger.Must(err)

	_, err = db.Collection(cachedLists).Indexes().CreateMany(ctx, []m.IndexModel{{
		Keys: bson.D{{
			Key:   "k",
			Value: 1,
		}, {
			Key:   "m",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	}, {
		Keys: bson.M{
			"ca": 1,
		},
		Options: options.Index().SetExpireAfterSeconds(int32((3 * 24 * time.Hour).Seconds())),
	}})
	logger.Must(err)

	_, err = db.Collection(reviews).Indexes().CreateMany(ctx, []m.IndexModel{{
		Keys: bson.D{{
			Key:   "c",
			Value: 1,
		}, {
			Key:   "s",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "u",
			Value: 1,
		}, {
			Key:   "s",
			Value: 1,
		}},
	}})
	logger.Must(err)

	_, err = db.Collection(categories).Indexes().CreateOne(context.Background(), m.IndexModel{
		Keys: bson.M{
			"s": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	logger.Must(err)

	_, err = db.Collection(cities).Indexes().CreateOne(context.Background(), m.IndexModel{
		Keys: bson.M{
			"s": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	logger.Must(err)

	_, err = db.Collection(technologies).Indexes().CreateOne(ctx, m.IndexModel{
		Keys: bson.M{
			"s": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	logger.Must(err)

	_, err = db.Collection(technologyCategories).Indexes().CreateOne(ctx, m.IndexModel{
		Keys: bson.M{
			"w": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	logger.Must(err)

	_, err = db.Collection(dns).Indexes().CreateOne(ctx, m.IndexModel{
		Keys: bson.M{
			"n": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	logger.Must(err)
}
