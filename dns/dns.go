package dns

import (
	"context"
	"encoding/json"
	"errors"
	rd "github.com/go-redis/redis/v8"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type DNS struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"n,omitempty" json:"n,omitempty"`
}

func (DNS) Create(
	ctx context.Context,
	name string,
) (
	res DNS,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = mongo.DNS.UpdateOne(ctx, DNS{
		Name: name,
	}, bson.M{
		"$setOnInsert": DNS{
			Name: name,
		},
	}, options.Update().SetUpsert(true))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	err = mongo.DNS.FindOne(ctx, DNS{
		Name: name,
	}).Decode(&res)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return
}

func (DNS) GetByIDs(ctx context.Context, ids []string) (res []DNS, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var (
		egGet     errgroup.Group
		mu        sync.Mutex
		cachedIds = make([]bool, len(ids))
		cachedDNS []DNS
	)
	for _i, _id := range ids {
		var (
			i  = _i
			id = _id
		)
		egGet.Go(func() (e error) {
			cacheVal, e := redis.Client.Get(ctx, makeDNSCacheKey(id)).Bytes()
			if e != nil {
				if errors.Is(e, rd.Nil) {
					e = nil
				}
				return
			}

			var dns DNS
			e = json.Unmarshal(cacheVal, &dns)
			if e != nil {
				return
			}

			mu.Lock()
			cachedIds[i] = true
			cachedDNS = append(cachedDNS, dns)
			mu.Unlock()
			return
		})
	}
	err = egGet.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = append(res, cachedDNS...)

	var oIDs []primitive.ObjectID
	for i, id := range ids {
		if cachedIds[i] {
			continue
		}

		oID, e := primitive.ObjectIDFromHex(id)
		if e != nil {
			err = e
			logger.Log.Error().Err(e).Send()
			return
		}
		oIDs = append(oIDs, oID)
	}
	if len(oIDs) == 0 {
		logger.Log.Debug().Strs("ids", ids).Msg("full cache")
		return
	}

	cur, err := mongo.DNS.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": oIDs,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var mongoDNS []DNS
	err = cur.All(ctx, &mongoDNS)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var egSet errgroup.Group
	for _, _md := range mongoDNS {
		md := _md
		egSet.Go(func() (e error) {
			cacheVal, e := json.Marshal(md)
			if e != nil {
				return
			}

			return redis.Client.Set(ctx, makeDNSCacheKey(md.ID.Hex()), cacheVal, 30*time.Second).Err()
		})
	}
	err = egSet.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = append(res, mongoDNS...)
	return
}

func (DNS) GetHints(ctx context.Context, name string, excludeIDs []string, limit int64) (
	res []DNS,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := bson.M{
		"n": bson.M{
			"$regex":   name,
			"$options": "$i",
		},
	}

	var oIDs []primitive.ObjectID
	for _, id := range excludeIDs {
		oID, e := primitive.ObjectIDFromHex(id)
		if e != nil {
			err = e
			logger.Log.Error().Err(e).Send()
			return
		}
		oIDs = append(oIDs, oID)
	}

	if len(oIDs) != 0 {
		query["_id"] = bson.M{
			"$nin": oIDs,
		}
	}

	cur, err := mongo.DNS.Find(ctx, query, options.Find().SetLimit(limit))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	err = cur.All(ctx, &res)
	logger.Err(err)
	return
}
