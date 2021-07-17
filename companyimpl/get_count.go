package companyimpl

import (
	"context"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
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

	count, err := mongo.Companies.CountDocuments(ctx, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetCountResponse{
		Count: uint32(count),
	}
	return
}
