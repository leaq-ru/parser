package technology

import (
	"context"
	"encoding/json"
	"errors"
	rd "github.com/go-redis/redis/v8"
	"github.com/gosimple/slug"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/parser/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type Technology struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	CategoryIDs []primitive.ObjectID `bson:"c,omitempty" json:"c,omitempty"`
	Name        string               `bson:"n,omitempty" json:"n,omitempty"`
	Slug        string               `bson:"s,omitempty" json:"s,omitempty"`
}

func (Technology) Create(
	ctx context.Context,
	wappalyzerTechnologyName string,
	categoryIDs []primitive.ObjectID,
) (
	res Technology,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if wappalyzerTechnologyName == "Cart Functionality" {
		wappalyzerTechnologyName = "Интернет-магазин"
	}

	techSlug := slug.Make(wappalyzerTechnologyName)

	_, err = mongo.Technologies.UpdateOne(ctx, Technology{
		Slug: techSlug,
	}, bson.M{
		"$setOnInsert": Technology{
			Name:        wappalyzerTechnologyName,
			CategoryIDs: categoryIDs,
		},
	}, options.Update().SetUpsert(true))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	err = mongo.Technologies.FindOne(ctx, Technology{
		Slug: techSlug,
	}).Decode(&res)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return
}

func (Technology) GetByIDs(ctx context.Context, ids []string) (res []Technology, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var (
		egGet       errgroup.Group
		mu          sync.Mutex
		cachedIds   = make([]bool, len(ids))
		cachedTechs []Technology
	)
	for _i, _id := range ids {
		var (
			i  = _i
			id = _id
		)
		egGet.Go(func() (e error) {
			cacheVal, e := redis.Client.Get(ctx, makeTechnologyCacheKey(id)).Bytes()
			if e != nil {
				if errors.Is(e, rd.Nil) {
					e = nil
				}
				return
			}

			var tech Technology
			e = json.Unmarshal(cacheVal, &tech)
			if e != nil {
				return
			}

			mu.Lock()
			cachedIds[i] = true
			cachedTechs = append(cachedTechs, tech)
			mu.Unlock()
			return
		})
	}
	err = egGet.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = append(res, cachedTechs...)

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

	cur, err := mongo.Technologies.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": oIDs,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var mongoTechs []Technology
	err = cur.All(ctx, &mongoTechs)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var egSet errgroup.Group
	for _, _mt := range mongoTechs {
		mt := _mt
		egSet.Go(func() (e error) {
			cacheVal, e := json.Marshal(mt)
			if e != nil {
				return
			}

			return redis.Client.Set(ctx, makeTechnologyCacheKey(mt.ID.Hex()), cacheVal, 30*time.Second).Err()
		})
	}
	err = egSet.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = append(res, mongoTechs...)
	return
}

func (Technology) GetHints(ctx context.Context, name string, excludeIDs []string, limit int64) (
	res []Technology,
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

	cur, err := mongo.Technologies.Find(ctx, query, options.Find().SetLimit(limit))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	err = cur.All(ctx, &res)
	logger.Err(err)
	return
}

func (Technology) GetBySlug(ctx context.Context, slug string) (
	res Technology,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = mongo.Technologies.FindOne(ctx, Technology{
		Slug: slug,
	}).Decode(&res)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = errors.New("technology not found")
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}
	return
}
