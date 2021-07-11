package company

import (
	"context"
	"github.com/nnqq/scr-parser/categoryimpl"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Company) setCategoryID(ctx context.Context, html string) {
	resCategory, err := categoryimpl.NewServer().FindCategory(ctx, &parser.FindCategoryRequest{
		Html: html,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	oID, err := primitive.ObjectIDFromHex(resCategory.GetCategoryId())
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	c.CategoryID = oID
}
