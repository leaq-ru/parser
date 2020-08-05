package companyimpl

import (
	"context"
	"errors"
	"fmt"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"time"
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

	const emailFieldName = "e"

	// force hasEmail=true
	for i, q := range query {
		if q.Key == emailFieldName {
			query[i] = makeExists(emailFieldName)
		}
	}

	emails, err := mongo.Companies.Distinct(ctx, emailFieldName, query)
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
