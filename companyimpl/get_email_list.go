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

const (
	email = "e"
	phone = "p"
)

func (s *server) GetEmailList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetEmailListResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	for i, q := range query {
		if q.Key == email {
			// force hasEmail=yes
			query[i] = bsonENotNil(email)
		}

		if q.Key == phone {
			// force hasPhone=any
			query[i] = bson.E{}
		}
	}

	emails, err := mongo.Companies.Distinct(ctx, email, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetEmailListResponse{}
	for _, e := range emails {
		email, ok := e.(string)
		if !ok {
			logger.Log.Error().Err(errors.New(fmt.Sprintf("can't process value email=%s", e))).Send()
			continue
		}

		res.Emails = append(res.Emails, email)
	}
	return
}
