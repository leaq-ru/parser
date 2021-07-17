package company

import (
	"context"
	"github.com/leaq-ru/parser/categoryimpl"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/proto/codegen/go/parser"
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
