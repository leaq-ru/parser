package companyimpl

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/company"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/md"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"github.com/nnqq/scr-proto/codegen/go/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (*server) Edit(ctx context.Context, req *parser.EditRequest) (
	res *empty.Empty,
	err error,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if req.GetCompanyId() == "" {
		err = errors.New("companyId required")
		return
	}

	authUserID, err := md.GetUserID(ctx)
	if err != nil {
		return
	}

	resRole, err := call.Role.CanEditCompany(ctx, &user.CanEditCompanyRequest{
		CompanyId: req.GetCompanyId(),
		UserId:    authUserID,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if !resRole.GetCanEdit() {
		err = errors.New("unauthorized")
		return
	}

	compOID, err := primitive.ObjectIDFromHex(req.GetCompanyId())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var categoryOID primitive.ObjectID
	if req.GetCategoryId() != nil {
		categoryOID, err = primitive.ObjectIDFromHex(req.GetCategoryId().GetValue())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	mongo.Companies.UpdateOne(ctx, company.Company{
		ID: compOID,
	}, bson.M{
		"$set": company.Company{
			Title:       req.GetTitle().GetValue(),
			Description: req.GetDescription().GetValue(),
			Email:       req.GetEmail().GetValue(),
			Phone:       req.GetPhone().GetValue(),
		},
	})
}
