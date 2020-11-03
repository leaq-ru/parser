package companyimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
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
