package companyimpl

import (
	"context"
	"errors"
	"fmt"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (s *server) GetPhoneList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetPhoneListResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	for i, q := range query {
		if q.Key == email {
			// force hasEmail=any
			query[i] = bson.E{}
		}

		if q.Key == phone {
			// force hasPhone=yes
			query[i] = makeExists(phone)
		}
	}

	phones, err := mongo.Companies.Distinct(ctx, phone, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetPhoneListResponse{}
	for _, p := range phones {
		switch phone := p.(type) {
		case int32:
			res.Phones = append(res.Phones, float64(phone))
		case int64:
			res.Phones = append(res.Phones, float64(phone))
		default:
			logger.Log.Error().Err(errors.New(fmt.Sprintf("can't process value phone=%s", p))).Send()
		}
	}
	return
}
