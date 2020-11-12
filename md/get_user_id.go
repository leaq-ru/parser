package md

import (
	"context"
	"errors"
	"github.com/nnqq/scr-parser/logger"
	"google.golang.org/grpc/metadata"
)

func GetUserID(ctx context.Context) (userID string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.New("failed to get metadata")
		logger.Log.Error().Err(err).Send()
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
