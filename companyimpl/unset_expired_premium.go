package companyimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (s *server) UnsetExpiredPremium(ctx context.Context, _ *empty.Empty) (res *empty.Empty, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = mongo.Companies.UpdateMany(ctx, bson.M{
		"pd": bson.M{
			"$lte": time.Now().UTC(),
		},
	}, bson.M{
		"$unset": bson.M{
			"pr": "",
			"pd": "",
		},
	})
	if err != nil {
		return
	}

	res = &empty.Empty{}
	return
}
