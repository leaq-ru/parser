package companyimpl

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/md"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/opts"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type opter interface {
	GetOpts() *opts.SkipLimit
}

func applyDefaultLimit(req opter) (limit int64, err error) {
	limit = int64(20)
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 || req.GetOpts().GetLimit() < 0 {
			err = errors.New("limit out of 1-100")
		} else if req.GetOpts().GetLimit() != 0 {
			limit = int64(req.GetOpts().GetLimit())
		}
	}
	return
}

func (*server) GetMy(ctx context.Context, req *parser.GetMyRequest) (
	res *parser.GetMyResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	limit, err := applyDefaultLimit(req)
	if err != nil {
		return
	}

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	resOwn, err := call.Role.GetOwnCompanies(ctx, &user.GetOwnCompaniesRequest{
		Opts: &opts.SkipLimit{
			Skip:  req.GetOpts().GetSkip(),
			Limit: uint32(limit),
		},
		UserId: authUserID,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var companyOIDs []primitive.ObjectID
	for _, id := range resOwn.GetCompanyIds() {
		oID, e := primitive.ObjectIDFromHex(id)
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}
		companyOIDs = append(companyOIDs, oID)
	}

	if len(companyOIDs) == 0 {
		res = &parser.GetMyResponse{}
		return
	}

	cur, err := mongo.companies.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": companyOIDs,
		},
	}, options.Find().SetSort(bson.M{
		"_id": -1,
	}))
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

	return fetchMyCompanies(ctx, comps)
}
