package companyimpl

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (*server) SetHidden(ctx context.Context, req *parser.SetHiddenRequest) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	in := make(bson.A, len(req.GetSlugs()))
	for i, s := range req.GetSlugs() {
		in[i] = s
	}

	_, err = mongo.companies.UpdateMany(ctx, bson.M{
		"s": bson.M{
			"$in": in,
		},
	}, bson.M{
		"$set": bson.M{
			"h": true,
		},
	})
	res = &empty.Empty{}
	return
}
