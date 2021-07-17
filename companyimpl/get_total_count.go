package companyimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"time"
)

func (*server) GetTotalCount(ctx context.Context, _ *empty.Empty) (res *parser.GetTotalCountResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	count, err := mongo.Companies.EstimatedDocumentCount(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetTotalCountResponse{
		TotalCount: uint32(count),
	}
	return
}
