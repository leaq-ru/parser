package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
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
		now := time.Now().UTC()

		_, e = mongo.Companies.UpdateOne(sc, bson.M{
			"$or": bson.A{bson.M{
				"_id": companyOID,
				"pd":  nil,
			}, bson.M{
				"_id": companyOID,
				"pd": bson.M{
					"$lt": now,
				},
			}},
		}, bson.M{
			"$set": company.Company{
				PremiumDeadline: now,
			},
		})
		if e != nil {
			return
		}

		monthAmount := time.Duration(req.GetMonthAmount()) * 31 * 24 * time.Hour

		_, e = mongo.Companies.UpdateOne(sc, company.Company{
			ID: companyOID,
		}, bson.A{bson.M{
			"$set": bson.M{
				"pd": bson.M{
					"$add": bson.A{
						"$pd",
						monthAmount.Milliseconds(),
					},
				},
				"pr": true,
				"ua": now,
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
