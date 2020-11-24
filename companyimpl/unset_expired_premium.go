package companyimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (s *server) UnsetExpiredPremium(ctx context.Context, _ *empty.Empty) (res *empty.Empty, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err = mongo.Companies.UpdateMany(ctx, bson.M{
		"pd": bson.M{
			"$lt": time.Now().UTC(),
		},
	}, bson.M{
		"$set": company.Company{
			UpdatedAt: time.Now().UTC(),
		},
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
