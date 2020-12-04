package companyimpl

import (
	"bytes"
	"context"
	"errors"
	"github.com/google/uuid"
	m "github.com/minio/minio-go/v7"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/md"
	"github.com/nnqq/scr-parser/minio"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

const (
	email = "e"
	phone = "p"
)

const freeListLimit = 1000

func (s *server) GetEmailList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetListResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	query, err := makeGetQuery(req)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	for i, q := range query {
		if q.Key == email {
			// force hasEmail=yes
			query[i] = bsonENotNil(email)
		}

		if q.Key == phone {
			// force hasPhone=any
			query[i] = bson.E{}
		}
	}

	premium, err := md.GetDataPremium(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		err = errors.New(http.StatusText(http.StatusInternalServerError))
		return
	}

	var opts *options.FindOptions
	if !premium {
		opts = options.Find().SetLimit(freeListLimit)
	}

	cur, err := mongo.Companies.Find(ctx, query, opts)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	defer func() {
		logger.Err(cur.Close(ctx))
	}()

	type onlyEmail struct {
		Email string `bson:"e"`
	}

	uniqueEmails := map[string]struct{}{}
	for cur.Next(ctx) {
		doc := onlyEmail{}
		e := cur.Decode(&doc)
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}

		if doc.Email != "" {
			uniqueEmails[doc.Email] = struct{}{}
		}
	}

	var file []byte
	for u := range uniqueEmails {
		file = append(file, []byte(u+"\n")...)
	}

	obj, err := minio.Client.PutObject(
		ctx,
		config.Env.S3.DownloadBucketName,
		"emails-"+uuid.New().String()+".txt",
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
