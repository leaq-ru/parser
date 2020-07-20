package company

import (
	"context"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-proto/codegen/go/image"
)

func (c *Company) setAvatar(ctx context.Context, url link) (err error) {
	s3res, err := call.Image.Put(ctx, &image.PutRequest{
		Url: string(url),
	})
	if err != nil {
		// debug level: often can't dl image due to CORS policy
		logger.Log.Debug().Str("url", string(url)).Err(err).Send()
		return
	}

	c.Avatar = link(s3res.GetS3Url())
	return
}
