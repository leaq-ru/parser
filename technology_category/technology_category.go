package technology_category

import (
	"context"
	"encoding/json"
	"errors"
	rd "github.com/go-redis/redis/v8"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/redis"
	"github.com/nnqq/scr-parser/translate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

type TechnologyCategory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"n,omitempty"`
	WappalyzerID uint32             `bson:"w,omitempty"`
}

func (TechnologyCategory) Create(ctx context.Context, wappalyzerCategoryID uint32) (res TechnologyCategory, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	name, ok := translate.TechnologyCategoryRU(wappalyzerCategoryID)
	if !ok {
		err = errors.New("invalid category")
		logger.Log.Error().Uint32("wappalyzerCategoryID", wappalyzerCategoryID).Err(err).Send()
		return
	}

	_, err = mongo.TechnologyCategories.UpdateOne(ctx, TechnologyCategory{
		WappalyzerID: wappalyzerCategoryID,
	}, bson.M{
		"$setOnInsert": TechnologyCategory{
			Name: name,
		},
	}, options.Update().SetUpsert(true))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	err = mongo.TechnologyCategories.FindOne(ctx, TechnologyCategory{
		WappalyzerID: wappalyzerCategoryID,
	}).Decode(&res)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	return
}

func (TechnologyCategory) GetByIDs(ctx context.Context, ids []string) (res []TechnologyCategory, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var (
		egGet      errgroup.Group
		mu         sync.Mutex
		cachedIds  = make([]bool, len(ids))
		cachedCats []TechnologyCategory
	)
	for _i, _id := range ids {
		var (
			i  = _i
			id = _id
		)
		egGet.Go(func() (e error) {
			cacheVal, e := redis.Client.Get(ctx, makeCategoryCacheKey(id)).Bytes()
			if e != nil {
				if errors.Is(e, rd.Nil) {
					e = nil
				}
				return
			}

			var cat TechnologyCategory
			e = json.Unmarshal(cacheVal, &cat)
			if e != nil {
				return
			}

			mu.Lock()
			cachedIds[i] = true
			cachedCats = append(cachedCats, cat)
			mu.Unlock()
			return
		})
	}
	err = egGet.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = append(res, cachedCats...)

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

	cur, err := mongo.TechnologyCategories.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": oIDs,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var mongoCats []TechnologyCategory
	err = cur.All(ctx, &mongoCats)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var egSet errgroup.Group
	for _, _mc := range mongoCats {
		mc := _mc
		egSet.Go(func() (e error) {
			cacheVal, e := json.Marshal(mc)
			if e != nil {
				return
			}

			return redis.Client.Set(ctx, makeCategoryCacheKey(mc.ID.Hex()), cacheVal, 30*time.Second).Err()
		})
	}
	err = egSet.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = append(res, mongoCats...)
	return
}
