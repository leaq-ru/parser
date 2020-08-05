package companyimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func makeExists(key string) bson.E {
	return bson.E{
		Key: key,
		Value: bson.M{
			"$exists": true,
		},
	}
}

func (s *server) GetRelated(ctx context.Context, req *parser.GetRelatedRequest) (
	res *parser.GetRelatedResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit := int64(5)
	if req.GetLimit() > 10 || req.GetLimit() < 0 {
		err = errors.New("limit out of 1-10")
		return
	} else if req.GetLimit() != 0 {
		limit = int64(req.GetLimit())
	}

	qCityCat, err := makeQueryCityCat(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	match := append(qCityCat, makeExists("e"), makeExists("p"))

	cur, err := mongo.Companies.Aggregate(ctx, []bson.M{
		{
			"$match": match,
		},
		{
			"$sample": bson.M{
				"size": limit,
			},
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var companies []model.Company
	err = cur.All(ctx, &companies)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	lenComps := int64(len(companies))
	if lenComps < limit {
		delta := limit - lenComps

		var excludeIDs []primitive.ObjectID
		for _, c := range companies {
			excludeIDs = append(excludeIDs, c.ID)
		}

		query := bson.M{
			"e": bson.M{
				"$exists": true,
			},
			"p": bson.M{
				"$exists": true,
			},
		}
		if len(excludeIDs) != 0 {
			query["_id"] = bson.M{
				"$nin": excludeIDs,
			}
		}

		opts := options.Find()
		opts.SetLimit(delta)
		cur, errExtraFind := mongo.Companies.Find(ctx, query, opts)
		if errExtraFind != nil {
			err = errExtraFind
			logger.Log.Error().Err(err).Send()
			return
		}

		var extraComps []model.Company
		err = cur.All(ctx, &extraComps)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		companies = append(companies, extraComps...)
	}

	var cityIDs, categoryIDs []string
	for _, c := range companies {
		if c.Location != nil && !c.Location.CityID.IsZero() {
			cityIDs = append(cityIDs, c.Location.CityID.Hex())
		}
		if !c.CategoryID.IsZero() {
			categoryIDs = append(categoryIDs, c.CategoryID.Hex())
		}
	}

	wg := sync.WaitGroup{}
	var (
		cities    *city.CitiesResponse
		errCities error
	)
	if len(cityIDs) != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cities, errCities = call.City.GetByIds(ctx, &city.GetByIdsRequest{
				CityIds: cityIDs,
			})
			logger.Err(errCities)
		}()
	}

	var (
		categories    *category.CategoriesResponse
		errCategories error
	)
	if len(categoryIDs) != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			categories, errCategories = call.Category.GetByIds(ctx, &category.GetByIdsRequest{
				CategoryIds: categoryIDs,
			})
			logger.Err(errCategories)
		}()
	}
	wg.Wait()

	if errCities != nil {
		err = errCities
		return
	}
	if errCategories != nil {
		err = errCategories
		return
	}

	res = &parser.GetRelatedResponse{}
	res.Companies, err = toFullCompanies(companies, cities, categories)
	return
}
