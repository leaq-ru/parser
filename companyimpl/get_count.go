package companyimpl

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"time"
)

func (s *server) GetCount(
	ctx context.Context,
	req *parser.GetV2Request,
) (
	res *parser.GetCountResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	query, err := makeGetQueryV2(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	count, err := mongo.companies.CountDocuments(ctx, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetCountResponse{
		Count: uint32(count),
	}
	return
}
