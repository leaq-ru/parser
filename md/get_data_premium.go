package md

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func GetDataPremium(ctx context.Context) (premium bool, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.New(http.StatusText(http.StatusInternalServerError))
		return
	}

	val := md.Get("data-premium")
	if len(val) != 0 {
		premium = val[0] == "true"
	}
	return
}
