package model

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/city"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Company) setCity(ctx context.Context, html string) {
	resCity, err := call.City.Find(ctx, &city.FindRequest{
		Html: html,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if resCity.IsFound {
		cityOid, err := primitive.ObjectIDFromHex(resCity.CityId)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		if c.Location == nil {
			c.Location = &location{}
		}
		c.Location.CityID = cityOid
	}
}
