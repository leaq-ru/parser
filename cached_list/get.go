package cached_list

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	m "go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/encoding/protojson"
	"strconv"
	"time"
)

func makeMD5Key(premium bool, key *parser.GetListRequest) (sum string, err error) {
	bytes, err := protojson.Marshal(key)
	if err != nil {
		return
	}

	rawSum := md5.Sum(append(bytes, []byte(strconv.FormatBool(premium))...))
	sum = hex.EncodeToString(rawSum[:])
	return
}

func Get(ctx context.Context, kind kind, premium bool, key *parser.GetListRequest) (s3URL string, cacheHit bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	md5Key, err := makeMD5Key(premium, key)
	if err != nil {
		return
	}

	var doc cachedList
	err = mongo.CachedLists.FindOne(ctx, cachedList{
		Kind: kind,
		MD5:  md5Key,
	}).Decode(&doc)
	if err != nil && errors.Is(err, m.ErrNoDocuments) {
		err = nil
		return
	}

	cacheHit = true
	s3URL = doc.URL
	return
}
