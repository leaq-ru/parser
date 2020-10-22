package company

import (
	"context"
	"errors"
	userAgent "github.com/EDDYCJY/fake-useragent"
	"github.com/nnqq/scr-parser/logger"
	"github.com/valyala/fasthttp"
	u "net/url"
)

func (c *Company) parseContactsPage(ctx context.Context) {
	parsedURL, err := u.Parse(c.URL)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	parsedURL.Path = "contacts"
	withSlug := parsedURL.String()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(withSlug)
	req.Header.SetUserAgent(userAgent.Random())
	res := fasthttp.AcquireResponse()
	err = makeSafeFastHTTPClient().DoRedirects(req, res, 3)
	if err != nil {
		if !errors.Is(err, fasthttp.ErrTooManyRedirects) {
			logger.Log.Error().Err(err).Send()
		}
		return
	}
	logger.Log.Debug().Str("withSlug", withSlug).Msg("blind related hit")

	var html []byte
	if enc := string(res.Header.Peek("Content-Encoding")); enc == "gzip" {
		body, err := res.BodyGunzip()
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		html = body
	} else {
		html = res.Body()
	}

	c.digHTML(ctx, html, false)
}
