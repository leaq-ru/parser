package companyimpl

import (
	"context"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (*server) GetManyBy(ctx context.Context, req *parser.GetManyByRequest) (
	res *parser.ShortCompanies,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	in := make(bson.A, len(req.GetCompanyIds()))
	for i, c := range req.GetCompanyIds() {
		oID, e := primitive.ObjectIDFromHex(c)
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}
		in[i] = oID
	}

	cur, err := mongo.Companies.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": in,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var comps []company.Company
	err = cur.All(ctx, &comps)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res, err = fetchShortCompanies(ctx, comps)
	return
}
