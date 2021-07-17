package company

import (
	"context"
	"github.com/leaq-ru/parser/call"
	"github.com/leaq-ru/parser/logger"
	"github.com/leaq-ru/proto/codegen/go/image"
)

func (c *Company) setAvatarWithUpload(ctx context.Context, rawURL Link) (err error) {
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
		c.Avatar = Link(s3res.GetS3Url())
	}
	return
}
