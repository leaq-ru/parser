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

func (s *server) GetPhoneList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetPhoneListResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	const phoneFieldName = "p"

	// force hasPhone=true
	for i, q := range query {
		if q.Key == phoneFieldName {
			query[i] = makeExists(phoneFieldName)
		}
	}

	phones, err := mongo.Companies.Distinct(ctx, phoneFieldName, query)
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
