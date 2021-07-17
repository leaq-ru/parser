package companyimpl

import (
	"context"
	"errors"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

func (s *server) GetBySlug(ctx context.Context, req *parser.GetBySlugRequest) (res *parser.FullCompany, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	comp := company.Company{}
	err = mongo.Companies.FindOne(ctx, bson.M{
		"s": req.GetSlug(),
	}).Decode(&comp)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			err = errors.New("company not found")
			return
		}

		logger.Log.Error().Err(err).Send()
		return
	}

	wg := sync.WaitGroup{}
	var (
		resCity *parser.CityItem
		errCity error
	)
	if comp.Location != nil && !comp.Location.CityID.IsZero() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resCity, errCity = cityimpl.NewServer().GetCityById(ctx, &parser.GetCityByIdRequest{
				CityId: comp.Location.CityID.Hex(),
			})
			if errCity != nil {
				logger.Log.Error().Err(errCity).Send()
			}
		}()
	}

	var (
		resCategory *parser.CategoryItem
		errCategory error
	)
	if !comp.CategoryID.IsZero() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resCategory, errCategory = categoryimpl.NewServer().GetCategoryById(ctx, &parser.GetCategoryByIdRequest{
				CategoryId: comp.CategoryID.Hex(),
			})
			if errCategory != nil {
				logger.Log.Error().Err(errCategory).Send()
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

	res = toFullCompany(comp, resCity, resCategory)
	return
}
