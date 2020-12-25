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
			Key:   "l.c",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "c",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "ti",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "o",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "he",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hp",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hv",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hi",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "ht",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hy",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hf",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "ha",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hg",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hin",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "hk",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "ho",
			Value: 1,
		}},
	}, {
		Keys: bson.D{{
			Key:   "so.v.m",
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
}
