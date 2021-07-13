package company

import (
	"context"
	"github.com/nnqq/scr-parser/dnsimpl"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (c *Company) withDNS(ctx context.Context, url string) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dnsIDs, err := dnsimpl.NewServer().FindDns(ctx, &parser.FindDnsRequest{
		Url: url,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	var dnsOIDs []primitive.ObjectID
	for _, id := range dnsIDs.GetIds() {
		oID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		dnsOIDs = append(dnsOIDs, oID)
	}

	c.DNSIDs = dnsOIDs
}
