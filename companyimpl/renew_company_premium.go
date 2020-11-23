package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func (s *server) RenewCompanyPremium(
	ctx context.Context,
	req *parser.RenewCompanyPremiumRequest,
) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetCompanyId() == "" || req.GetMonthAmount() == 0 {
		err = errors.New(http.StatusText(http.StatusBadRequest))
		return
	}

	companyOID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	sess, err := mongo.Client.StartSession()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	_, err = sess.WithTransaction(ctx, func(sc m.SessionContext) (_ interface{}, e error) {
		_, e = mongo.Companies.UpdateOne(sc, company.Company{
			ID:              companyOID,
			PremiumDeadline: nil,
		}, bson.M{
			"$set": company.Company{
				PremiumDeadline: time.Now().UTC(),
			},
		})
		if e != nil {
			return
		}

		_, e = mongo.Companies.UpdateOne(ctx, company.Company{
			ID: companyOID,
		}, bson.A{bson.M{
			"$set": bson.M{
				"pd": bson.M{
					"$add": bson.A{
						"$pd",
						time.Duration(req.GetMonthAmount()) * 31 * 24 * time.Hour,
					},
				},
				"pr": true,
			},
		}})
		return
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &empty.Empty{}
	return
}
