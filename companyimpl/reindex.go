package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/leaq-ru/parser/company"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"time"
)

func (s *server) Reindex(ctx context.Context, req *parser.ReindexRequest) (res *empty.Empty, err error) {
	return s.reindex(ctx, req, false)
}

func (s *server) reindex(ctx context.Context, req *parser.ReindexRequest, async bool) (res *empty.Empty, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if req.GetUrl() == "" {
		err = errors.New("url required")
		logger.Log.Error().Err(err).Send()
		return
	}

	var t time.Time
	if req.GetRegistrationDate().IsValid() {
		t, err = ptypes.Timestamp(req.GetRegistrationDate())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	comp := company.Company{}
	comp.UpdateOrCreate(ctx, req.GetUrl(), req.GetRegistrar(), t, async)

	res = &empty.Empty{}
	return
}
