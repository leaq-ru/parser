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

	_, err := db.Collection(companies).Indexes().CreateMany(ctx, []m.IndexModel{
		{
			Keys: bson.M{
				"u": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"s": 1,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{
				Key:   "l.c",
				Value: 1,
			}, {
				Key:   "pd",
				Value: -1,
			}, {
				Key:   "_id",
				Value: -1,
			}},
			Options: options.Index().SetPartialFilterExpression(bson.M{
				"l.c": bson.M{
					"$exists": true,
				},
			}),
		},
		{
			Keys: bson.D{{
				Key:   "c",
				Value: 1,
			}, {
				Key:   "pd",
				Value: -1,
			}, {
				Key:   "_id",
				Value: -1,
			}},
			Options: options.Index().SetPartialFilterExpression(bson.M{
				"c": bson.M{
					"$exists": true,
				},
			}),
		},
		{
			Keys: bson.M{
				"e": 1,
			},
		},
		{
			Keys: bson.M{
				"p": 1,
			},
		},
		{
			Keys: bson.M{
				"o": 1,
			},
		},
		{
			Keys: bson.M{
				"i": 1,
			},
		},
		{
			Keys: bson.M{
				"k": 1,
			},
		},
		{
			Keys: bson.M{
				"og": 1,
			},
		},
		{
			Keys: bson.M{
				"ap.a.u": 1,
			},
		},
		{
			Keys: bson.M{
				"ap.g.u": 1,
			},
		},
		{
			Keys: bson.M{
				"so.v.g": 1,
			},
		},
		{
			Keys: bson.M{
				"so.v.m": 1,
			},
		},
		{
			Keys: bson.M{
				"so.i.u": 1,
			},
		},
		{
			Keys: bson.M{
				"so.t.u": 1,
			},
		},
		{
			Keys: bson.M{
				"so.y.u": 1,
			},
		},
		{
			Keys: bson.M{
				"so.f.u": 1,
			},
		},
		{
			Keys: bson.D{{
				Key:   "c",
				Value: 1,
			}, {
				Key:   "l.c",
				Value: 1,
			}, {
				Key:   "pd",
				Value: -1,
			}, {
				Key:   "_id",
				Value: -1,
			}},
			Options: options.Index().SetPartialFilterExpression(bson.D{{
				Key: "c",
				Value: bson.M{
					"$exists": true,
				},
			}, {
				Key: "l.c",
				Value: bson.M{
					"$exists": true,
				},
			}}),
		},
		{
			Keys: bson.D{{
				Key:   "ti",
				Value: 1,
			}, {
				Key:   "pd",
				Value: -1,
			}, {
				Key:   "_id",
				Value: -1,
			}},
			Options: options.Index().SetPartialFilterExpression(bson.M{
				"ti": bson.M{
					"$exists": true,
				},
			}),
		},
		{
			Keys: bson.D{{
				Key:   "pd",
				Value: -1,
			}, {
				Key:   "_id",
				Value: -1,
			}},
		},
	})
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
}
