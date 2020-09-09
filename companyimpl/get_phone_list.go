package companyimpl

import (
	"context"
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
			query[i] = bsonENotNil(phone)
		}
	}

	cur, err := mongo.Companies.Find(ctx, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	type onlyPhone struct {
		Phone float64 `bson:"p"`
	}

	uniquePhones := map[float64]struct{}{}
	for cur.Next(ctx) {
		if len(uniquePhones) >= limitListDownload {
			break
		}

		doc := onlyPhone{}
		err = cur.Decode(&doc)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		if doc.Phone != 0 {
			uniquePhones[doc.Phone] = struct{}{}
		}
	}

	res = &parser.GetPhoneListResponse{}
	for p := range uniquePhones {
		res.Phones = append(res.Phones, p)
	}
	return
}
