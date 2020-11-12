package companyimpl

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/category"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/sync/errgroup"
	"time"
)

func (*server) GetBy(ctx context.Context, req *parser.GetByRequest) (
	res *parser.ShortCompany,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := bson.M{}
	if req.GetUrl() != "" {
		query["u"] = req.GetUrl()
	}

	var comp company.Company
	err = mongo.Companies.FindOne(ctx, query).Decode(&comp)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var eg errgroup.Group
	var cityItem *city.CityItem
	if comp.Location != nil && !comp.Location.CityID.IsZero() {
		eg.Go(func() error {
			cityItem, err = call.City.GetById(ctx, &city.GetByIdRequest{
				CityId: comp.Location.CityID.Hex(),
			})
			return err
		})
	}

	var categoryItem *category.CategoryItem
	if !comp.CategoryID.IsZero() {
		eg.Go(func() error {
			categoryItem, err = call.Category.GetById(ctx, &category.GetByIdRequest{
				CategoryId: comp.CategoryID.Hex(),
			})
			return err
		})
	}
	err = eg.Wait()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = toShortCompany(comp, cityItem, categoryItem)
	return
}
