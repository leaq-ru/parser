package companyimpl

import (
	"context"
	"errors"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/parser/technologyimpl"
	"github.com/leaq-ru/proto/codegen/go/parser"
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

			var resCity *parser.CityItem
			resCity, errCity = cityimpl.NewServer().GetCityBySlug(ctx, &parser.GetCityBySlugRequest{
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

			var resCategory *parser.CategoryItem
			resCategory, errCategory = categoryimpl.NewServer().GetCategoryBySlug(ctx, &parser.GetCategoryBySlugRequest{
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

			var resTech *parser.GetTechBySlugResponse
			resTech, errTech = technologyimpl.NewServer().GetTechBySlug(ctx, &parser.GetTechBySlugRequest{
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

	cur, err := mongo.Companies.Find(ctx, query, opts)
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
		cities    *parser.CitiesResponse
		errCities error
	)
	if len(cityIDs) != 0 {
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
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
		wgFullDocs.Add(1)
		go func() {
			defer wgFullDocs.Done()
			categories, errCategories = categoryimpl.NewServer().GetCategoryByIds(ctx, &parser.GetCategoryByIdsRequest{
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
