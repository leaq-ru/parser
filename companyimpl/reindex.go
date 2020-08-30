package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/model"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"time"
)

func (s *server) Reindex(ctx context.Context, req *parser.ReindexRequest) (res *empty.Empty, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	if req.GetUrl() == "" {
		err = errors.New("url required")
		logger.Log.Error().Err(err).Send()
		return
	}

	t, err := ptypes.Timestamp(req.GetRegistrationDate())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	comp := model.Company{}
	comp.UpdateOrCreate(ctx, req.GetUrl(), req.GetRegistrar(), t)

	res = &empty.Empty{}
	return
}
