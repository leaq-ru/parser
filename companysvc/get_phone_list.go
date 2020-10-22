package companysvc

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	m "github.com/minio/minio-go/v7"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/minio"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

func (s *server) GetPhoneList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetListResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	for i, q := range query {
		if q.Key == email {
			// force hasEmail=any
			query[i] = bson.E{}
		}

		if q.Key == phone {
			// force hasPhone=yes
			query[i] = bsonENotNil(phone)
		}
	}

	cur, err := mongo.Companies.Find(ctx, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	defer func() {
		logger.Err(cur.Close(ctx))
	}()

	type onlyPhone struct {
		Phone int `bson:"p"`
	}

	uniquePhones := map[int]struct{}{}
	for cur.Next(ctx) {
		doc := onlyPhone{}
		err = cur.Decode(&doc)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		if doc.Phone != 0 {
			uniquePhones[doc.Phone] = struct{}{}
		}
	}

	var file []byte
	for u := range uniquePhones {
		file = append(file, []byte(strconv.Itoa(u)+"\n")...)
	}

	obj, err := minio.Client.PutObject(
		ctx,
		config.Env.S3.DownloadBucketName,
		"phones-"+uuid.New().String()+".txt",
		bytes.NewReader(file),
		int64(len(file)),
		m.PutObjectOptions{},
	)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.GetListResponse{
		DownloadUrl: "https://" + config.Env.S3.DownloadBucketName + ".ru/" + obj.Key,
	}
	return
}
