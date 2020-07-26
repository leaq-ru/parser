package companyimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func (s *server) Get(ctx context.Context, req *parser.GetRequest) (res *parser.GetResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	skip := int64(0)
	limit := int64(20)
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 || req.GetOpts().GetLimit() < 1 {
			err = errors.New("limit out of 1-100")
		}
		if req.GetOpts().GetSkip() < 0 {
			err = errors.New("skip less than 0")
		}

		skip = int64(req.GetOpts().GetSkip())
		limit = int64(req.GetOpts().GetLimit())
	}

	query := bson.D{}
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
		if err != nil {
			err = errOID
			logger.Log.Error().Err(err).Send()
			return
		}

		query = append(query, bson.E{
			Key:   "c",
			Value: oID,
		})
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	var (
		comps   []model.Company
		errFind error
	)
	go func() {
		defer wg.Done()
		opts := options.Find()
		opts.SetSkip(skip)
		opts.SetLimit(limit)
		opts.SetSort(bson.M{
			"_id": 1,
		})
		opts.SetProjection(bson.M{
			"_id": 1,
			"c":   1,
			"l":   1,
			"u":   1,
			"s":   1,
			"t":   1,
			"e":   1,
			"o":   1,
			"p":   1,
			"a":   1,
		})
		var cur *m.Cursor
		cur, errFind = mongo.Companies.Find(ctx, query, opts)
		if errFind != nil {
			logger.Log.Error().Err(errFind).Send()
			return
		}
		errFind = cur.All(ctx, &comps)
		if errFind != nil {
			logger.Log.Error().Err(errFind).Send()
		}
	}()

	var (
		totalCount int64
		errCount   error
	)
	go func() {
		defer wg.Done()
		totalCount, errCount = mongo.Companies.CountDocuments(ctx, query)
		if errCount != nil {
			logger.Log.Error().Err(errCount).Send()
		}
	}()
	wg.Wait()

	if errFind != nil {
		err = errFind
		return
	}
	if errCount != nil {
		err = errCount
		return
	}

	res = &parser.GetResponse{
		TotalCount: uint32(totalCount),
	}
	for _, c := range comps {
		var location *parser.ShortLocation
		if c.Location != nil {
			location = &parser.ShortLocation{}
			location.AddressTitle = c.Location.AddressTitle
			location.Address = c.Location.Address
			location.CityId = c.Location.CityID.Hex()
		}

		categoryID := ""
		if !c.CategoryID.IsZero() {
			categoryID = c.CategoryID.Hex()
		}

		res.ShortCompanies = append(res.ShortCompanies, &parser.ShortCompany{
			Id:         c.ID.Hex(),
			CategoryId: categoryID,
			Location:   location,
			Url:        c.URL,
			Slug:       c.Slug,
			Title:      c.Title,
			Email:      c.Email,
			Online:     c.Online,
			Phone:      float64(c.Phone),
			Avatar:     string(c.Avatar),
		})
	}
	return
}
