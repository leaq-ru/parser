package mongo

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndex(db *mongo.Database) {
	companies := db.Collection(companies)
	_, err := companies.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
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
			Keys: bson.M{
				"l.c": 1,
			},
			Options: options.Index().SetPartialFilterExpression(bson.M{
				"l.c": bson.M{
					"$exists": true,
				},
			}),
		},
		{
			Keys: bson.M{
				"c": 1,
			},
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
			Keys: bson.M{
				"a": 1,
			},
		},
		{
			Keys: bson.M{
				"c":   1,
				"l.c": 1,
			},
			Options: options.Index().SetPartialFilterExpression(bson.M{
				"c": bson.M{
					"$exists": true,
				},
				"l.c": bson.M{
					"$exists": true,
				},
			}),
		},
	})
	logger.Must(err)
}
