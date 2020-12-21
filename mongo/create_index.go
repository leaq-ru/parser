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

	idxToDrop := []string{
		"e_1",
		"p_1",
		"o_1",
		"i_1",
		"k_1",
		"og_1",
		"ap.a.u_1",
		"ap.g.u_1",
		"so.v.g_1",
		"so.v.m_1",
		"so.i.u_1",
		"so.t.u_1",
		"so.y.u_1",
		"so.f.u_1",
	}
	for _, idx := range idxToDrop {
		_, err := db.Collection(companies).Indexes().DropOne(ctx, idx)
		logger.Err(err)
	}

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
			Keys: bson.D{{
				Key:   "o",
				Value: 1,
			}, {
				Key:   "he",
				Value: 1,
			}, {
				Key:   "hp",
				Value: 1,
			}, {
				Key:   "hv",
				Value: 1,
			}, {
				Key:   "hi",
				Value: 1,
			}, {
				Key:   "ht",
				Value: 1,
			}, {
				Key:   "hy",
				Value: 1,
			}, {
				Key:   "hf",
				Value: 1,
			}, {
				Key:   "ha",
				Value: 1,
			}, {
				Key:   "hg",
				Value: 1,
			}, {
				Key:   "hin",
				Value: 1,
			}, {
				Key:   "hk",
				Value: 1,
			}, {
				Key:   "ho",
				Value: 1,
			}, {
				Key:   "pd",
				Value: -1,
			}, {
				Key:   "_id",
				Value: -1,
			}, {
				Key:   "l.c",
				Value: 1,
			}, {
				Key:   "c",
				Value: 1,
			}, {
				Key:   "ti",
				Value: 1,
			}, {
				Key:   "so.v.m",
				Value: 1,
			}},
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
