package companyimpl

import (
	"bytes"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/leaq-ru/parser/cached_list"
	"github.com/leaq-ru/parser/config"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/md"
	"github.com/leaq-ru/parser/minio"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	m "github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

func (s *server) GetPhoneList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetListResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	ise := errors.New(http.StatusText(http.StatusInternalServerError))

	premium, err := md.GetDataPremium(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = ise
		return
	}

	cachedS3URL, cacheHit, err := cached_list.Get(ctx, cached_list.Kind_phone, premium, req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = ise
		return
	}

	if cacheHit {
		res = &parser.GetListResponse{
			DownloadUrl: cachedS3URL,
		}
		return
	}

	query, err := makeGetQueryV2(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = ise
		return
	}

	for i, q := range query {
		if q.Key == emailKey {
			// force email=any
			query[i] = bson.E{}
		}

		if q.Key == phoneKey {
			// force phone not nil
			query[i] = bsonENotNil(phoneKey)
		}
	}

	opts := options.Find()
	if !premium {
		opts.SetLimit(freeListLimit)
	}

	cur, err := mongo.Companies.Find(ctx, query, opts.SetProjection(bson.M{
		"_id": -1,
		"p":   1,
	}))
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
		err = ise
		return
	}

	s3URL := "https://" + config.Env.S3.DownloadBucketName + ".ru/" + obj.Key

	err = cached_list.Set(ctx, cached_list.Kind_phone, premium, req, s3URL)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = ise
		return
	}

	res = &parser.GetListResponse{
		DownloadUrl: s3URL,
	}
	return
}
