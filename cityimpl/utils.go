package cityimpl

import (
	"context"
	"github.com/leaq-ru/parser/city"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/proto/codegen/go/parser"
	m "go.mongodb.org/mongo-driver/mongo"
)

func citiesCursorToCitiesResponse(ctx context.Context, cur *m.Cursor) (res *parser.CitiesResponse, err error) {
	var cities []city.City
	err = cur.All(ctx, &cities)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	res = &parser.CitiesResponse{}
	for _, c := range cities {
		res.Cities = append(res.Cities, &parser.CityItem{
			Id:    c.ID.Hex(),
			Title: c.Title,
			Slug:  c.Slug,
		})
	}
	return
}
