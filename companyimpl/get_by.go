package companyimpl

import (
	"context"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if req.GetCompanyId() != "" {
		query["_id"], err = primitive.ObjectIDFromHex(req.GetCompanyId())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	var comp company.Company
	err = mongo.Companies.FindOne(ctx, query).Decode(&comp)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var eg errgroup.Group
	var cityItem *parser.CityItem
	if comp.Location != nil && !comp.Location.CityID.IsZero() {
		eg.Go(func() (e error) {
			cityItem, e = cityimpl.NewServer().GetCityById(ctx, &parser.GetCityByIdRequest{
				CityId: comp.Location.CityID.Hex(),
			})
			return
		})
	}

	var categoryItem *parser.CategoryItem
	if !comp.CategoryID.IsZero() {
		eg.Go(func() (e error) {
			categoryItem, e = categoryimpl.NewServer().GetCategoryById(ctx, &parser.GetCategoryByIdRequest{
				CategoryId: comp.CategoryID.Hex(),
			})
			return
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
