package company

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/technology"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (c *Company) withDNS(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	dnsIDs, err := call.DNS.FindDns(ctx, &technology.FindDnsRequest{
		Url: c.URL,
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
