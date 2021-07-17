package company

import (
	"context"
	"github.com/leaq-ru/parser/cityimpl"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Company) setCityID(ctx context.Context, html string) {
	resCity, err := cityimpl.NewServer().FindCity(ctx, &parser.FindCityRequest{
		Html: html,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if resCity.GetIsFound() {
		oID, err := primitive.ObjectIDFromHex(resCity.GetCityId())
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		if c.Location == nil {
			c.Location = &location{}
		}
		if c.Location.CityID.IsZero() {
			c.Location.CityID = oID
		}
	}
}
