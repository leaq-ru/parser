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
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

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
		{
			"$project": shortCompanyProjection,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var comps []model.Company
	err = cur.All(ctx, &comps)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	lenComps := int64(len(comps))
	if lenComps < limit {
		delta := limit - lenComps

		var excludeIDs []primitive.ObjectID
		for _, c := range comps {
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
		opts.SetProjection(shortCompanyProjection)
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
		comps = append(comps, extraComps...)
	}

	res = &parser.GetRelatedResponse{}
	for _, c := range comps {
		var location *parser.ShortLocation
		if c.Location != nil {
			cityID := ""
			if !c.Location.CityID.IsZero() {
				cityID = c.Location.CityID.Hex()
			}

			location = &parser.ShortLocation{
				CityId:       cityID,
				Address:      c.Location.Address,
				AddressTitle: c.Location.AddressTitle,
			}
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
