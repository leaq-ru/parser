package cached_list

import (
	"context"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Set(ctx context.Context, kind kind, premium bool, key *parser.GetListRequest, s3URL string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	md5Key, err := makeMD5Key(premium, key)
	if err != nil {
		return
	}

	_, err = mongo.CachedLists.UpdateOne(ctx, cachedList{
		Kind: kind,
		MD5:  md5Key,
	}, bson.M{
		"$set": cachedList{
			URL:       s3URL,
			CreatedAt: time.Now().UTC(),
		},
	}, options.Update().SetUpsert(true))
	return
}
