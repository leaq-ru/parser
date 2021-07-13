package companyimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/categoryimpl"
	"github.com/nnqq/scr-parser/cityimpl"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func bsonENotNil(key string) bson.E {
	return bson.E{
		Key: key,
		Value: bson.M{
			"$ne": nil,
		},
	}
}

func makeQuerySingleCityCat(req cityCat) (query bson.D, err error) {
	query = bson.D{}
	if req.GetCityId() != "" {
		oID, errOID := primitive.ObjectIDFromHex(req.GetCityId())
		if errOID != nil {
			err = errOID
			logger.Log.Error().Err(err).Send()
			return
		}

		query = append(query, bson.E{
			Key:   "l.c",
			Value: oID,
		})
	}
	if req.GetCategoryId() != "" {
		oID, errOID := primitive.ObjectIDFromHex(req.GetCategoryId())
		if errOID != nil {
			err = errOID
			logger.Log.Error().Err(err).Send()
			return
		}

		query = append(query, bson.E{
			Key:   "c",
			Value: oID,
		})
	}
	return
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

	qCityCat, err := makeQuerySingleCityCat(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	cur, err := mongo.Companies.Find(ctx, qCityCat, options.Find().SetLimit(limit))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var companies []company.Company
	err = cur.All(ctx, &companies)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
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
		cities    *parser.CitiesResponse
		errCities error
	)
	if len(cityIDs) != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cities, errCities = cityimpl.NewServer().GetCityByIds(ctx, &parser.GetCityByIdsRequest{
				CityIds: cityIDs,
			})
			logger.Err(errCities)
		}()
	}

	var (
		categories    *parser.CategoriesResponse
		errCategories error
	)
	if len(categoryIDs) != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			categories, errCategories = categoryimpl.NewServer().GetCategoryByIds(ctx, &parser.GetCategoryByIdsRequest{
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
