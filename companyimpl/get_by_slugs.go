package companyimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/technology"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func (s *server) GetBySlugs(ctx context.Context, req *parser.GetBySlugsRequest) (res *parser.ShortCompanies, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit := int64(20)
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 || req.GetOpts().GetLimit() < 0 {
			err = errors.New("limit out of 1-100")
			return
		} else if req.GetOpts().GetLimit() != 0 {
			limit = int64(req.GetOpts().GetLimit())
		}
	}

	query := bson.D{}
	if len(req.GetOpts().GetExcludeIds()) != 0 {
		var oIDs []primitive.ObjectID
		for _, id := range req.GetOpts().GetExcludeIds() {
			oID, e := primitive.ObjectIDFromHex(id)
			if e != nil {
				err = e
				logger.Log.Error().Err(err).Send()
				return
			}

			oIDs = append(oIDs, oID)
		}

		query = append(query, bson.E{
			Key: "_id",
			Value: bson.M{
				"$nin": oIDs,
			},
		})
	}

	var wg sync.WaitGroup
	var (
		cityID  primitive.ObjectID
		errCity error
	)
	if req.GetCitySlug() != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var resCity *city.CityItem
			resCity, errCity = call.City.GetBySlug(ctx, &city.GetBySlugRequest{
				Slug: req.GetCitySlug(),
			})
			if errCity != nil {
				logger.Log.Error().Err(errCity).Send()
				return
			}

			cityID, errCity = primitive.ObjectIDFromHex(resCity.GetId())
			if errCity != nil {
				logger.Log.Error().Err(errCity).Send()
			}
		}()
	}

	var (
		categoryID  primitive.ObjectID
		errCategory error
	)
	if req.GetCategorySlug() != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var resCategory *category.CategoryItem
			resCategory, errCategory = call.Category.GetBySlug(ctx, &category.GetBySlugRequest{
				Slug: req.GetCategorySlug(),
			})
			if errCategory != nil {
				logger.Log.Error().Err(errCategory).Send()
				return
			}

			categoryID, errCategory = primitive.ObjectIDFromHex(resCategory.GetId())
			if errCategory != nil {
				logger.Log.Error().Err(errCategory).Send()
			}
		}()
	}

	var (
		techID  primitive.ObjectID
		errTech error
	)
	if req.GetTechnologySlug() != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var resTech *technology.GetBySlugResponse
			resTech, errTech = call.Technology.GetBySlug(ctx, &technology.GetBySlugRequest{
				Slug: req.GetTechnologySlug(),
			})
			if errTech != nil {
				logger.Log.Error().Err(errTech).Send()
				return
			}

			techID, errTech = primitive.ObjectIDFromHex(resTech.GetTechnology().GetId())
			if errTech != nil {
				logger.Log.Error().Err(errTech).Send()
			}
		}()
	}
	wg.Wait()

	if errCity != nil {
		err = errCity
		return
	}
	if errCategory != nil {
		err = errCategory
		return
	}
	if errTech != nil {
		err = errTech
		return
	}

	if !cityID.IsZero() {
		query = append(query, bson.E{
			Key:   "l.c",
			Value: cityID,
		})
	}
	if !categoryID.IsZero() {
		query = append(query, bson.E{
			Key:   "c",
			Value: categoryID,
		})
	}
	if !techID.IsZero() {
		query = append(query, bson.E{
			Key:   "ti",
			Value: techID,
		})
	}

	opts := options.Find()
	opts.SetSkip(int64(req.GetOpts().GetSkip()))
	opts.SetLimit(limit)
	opts.SetSort(bson.M{
		"_id": -1,
	})

	cur, err := mongo.companies.Find(ctx, query, opts)
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

	wgFullDocs := sync.WaitGroup{}
	var (
		cities    *city.CitiesResponse
		errCities error
	)
	if len(cityIDs) != 0 {
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
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
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
			categories, errCategories = call.Category.GetByIds(ctx, &category.GetByIdsRequest{
				CategoryIds: categoryIDs,
			})
			logger.Err(errCategories)
		}()
	}
	wgFullDocs.Wait()

	if errCities != nil {
		err = errCities
		return
	}
	if errCategories != nil {
		err = errCategories
		return
	}

	res = &parser.ShortCompanies{}
	res.Companies, err = toShortCompanies(companies, cities, categories)
	return
}
