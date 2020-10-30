package company

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/image"
)

func (c *Company) setAvatar(ctx context.Context, rawURL link) (err error) {
	url := string(rawURL)

	s3res, err := call.Image.Put(ctx, &image.PutRequest{
		Url: url,
	})
	if err != nil {
		// debug level: often can't dl image due to CORS policy
		logger.Log.Debug().Str("url", url).Err(err).Send()
		return
	}

	if s3res.GetS3Url() != "" {
		c.Avatar = link(s3res.GetS3Url())
	}
	return
}
