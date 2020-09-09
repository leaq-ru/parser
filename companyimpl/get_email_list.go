package companyimpl

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	limitListDownload = 100000
	email             = "e"
	phone             = "p"
)

func (s *server) GetEmailList(ctx context.Context, req *parser.GetListRequest) (
	res *parser.GetEmailListResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
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

	cur, err := mongo.Companies.Find(ctx, query)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	type onlyEmail struct {
		Email string `bson:"e"`
	}

	uniqueEmails := map[string]struct{}{}
	for cur.Next(ctx) {
		if len(uniqueEmails) >= limitListDownload {
			break
		}

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

	res = &parser.GetEmailListResponse{}
	for e := range uniqueEmails {
		res.Emails = append(res.Emails, e)
	}
	return
}
