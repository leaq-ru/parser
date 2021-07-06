package md

import (
	"context"
	"errors"
	safeerr "github.com/nnqq/scr-lib-safeerr"
	"github.com/nnqq/scr-parser/logger"
	"google.golang.org/grpc/metadata"
)

func GetUserID(ctx context.Context) (userID string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = safeerr.InternalServerError
		return
	}

	val := md.Get("user-id")
	if len(val) != 0 {
		userID = val[0]
	}

	if userID == "" {
		err = errors.New("unauthorized")
		logger.Log.Error().Err(err).Send()
	}
	return
}
