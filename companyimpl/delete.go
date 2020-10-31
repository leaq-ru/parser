package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/post"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"time"
)

func deleteCompanyWithRefs(ctx context.Context, comp company.Company) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	sess, err := mongo.Client.StartSession()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	defer sess.EndSession(ctx)

	_, err = sess.WithTransaction(ctx, func(sc m.SessionContext) (_ interface{}, e error) {
		_, e = mongo.Companies.DeleteOne(sc, company.Company{
			ID: comp.ID,
		})
		if e != nil {
			logger.Log.Error().Err(e).Send()
			return
		}

		_, e = mongo.Posts.DeleteMany(sc, post.Post{
			CompanyID: comp.ID,
		})
		if e != nil {
			logger.Log.Error().Err(e).Send()
		}
		return
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()

		errAbort := sess.AbortTransaction(ctx)
		if errAbort != nil {
			err = errAbort
			logger.Log.Error().Err(err).Send()
		}
		return
	}

	err = sess.CommitTransaction(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if comp.Avatar != "" {
		_, err = call.Image.Remove(ctx, &image.RemoveRequest{
			S3Url: string(comp.Avatar),
		})
		if err != nil {
			logger.Log.Error().Err(err).Send()
		}
	}
	return
}

func (s *server) Delete(ctx context.Context, req *parser.DeleteRequest) (res *empty.Empty, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var oIDs []primitive.ObjectID
	for _, id := range req.GetIds() {
		oID, e := primitive.ObjectIDFromHex(id)
		if e != nil {
			err = e
			logger.Log.Error().Err(err).Send()
			return
		}

		oIDs = append(oIDs, oID)
	}

	if len(oIDs) == 0 {
		err = errors.New("at least one ID required")
		logger.Log.Error().Err(err).Send()
		return
	}

	cur, err := mongo.Companies.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": oIDs,
		},
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	for cur.Next(ctx) {
		var comp company.Company
		err = cur.Decode(&comp)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		err = deleteCompanyWithRefs(ctx, comp)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	res = &empty.Empty{}
	return
}
