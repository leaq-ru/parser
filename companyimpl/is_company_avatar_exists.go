package companyimpl

import (
	"context"
	"errors"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/parser/mongo"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (s *server) IsCompanyAvatarExists(ctx context.Context, req *parser.IsCompanyAvatarExistsRequest) (
	res *parser.IsCompanyAvatarExistsResponse,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res = &parser.IsCompanyAvatarExistsResponse{}

	errFindOne := mongo.Companies.FindOne(ctx, bson.M{
		"a": req.GetAvatarS3Url(),
	}).Err()
	if errFindOne != nil {
		if errors.Is(errFindOne, m.ErrNoDocuments) {
			res.IsExists = false
			return
		} else {
			err = errFindOne
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	res.IsExists = true
	return
}
