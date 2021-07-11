package categoryimpl

import (
	"context"
	"github.com/nnqq/scr-parser/category"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	m "go.mongodb.org/mongo-driver/mongo"
)

func categoriesCursorToCategoriesResponse(ctx context.Context, cur *m.Cursor) (
	res *parser.CategoriesResponse,
	err error,
) {
	var cats []category.Category
	err = cur.All(ctx, &cats)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.CategoriesResponse{}
	for _, c := range cats {
		res.Categories = append(res.Categories, &parser.CategoryItem{
			Id:    c.ID.Hex(),
			Title: string(c.Title),
			Slug:  c.Slug,
		})
	}
	return
}
