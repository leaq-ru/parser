package model

import (
	"context"
	"errors"
	userAgent "github.com/EDDYCJY/fake-useragent"
	"github.com/nnqq/scr-parser/logger"
	"github.com/valyala/fasthttp"
	u "net/url"
	"sync"
)

func (c *Company) parseRelatedPages(ctx context.Context, slugs ...string) {
	var (
		mu    sync.Mutex
		htmls [][]byte
	)

	wg := sync.WaitGroup{}
	wg.Add(len(slugs))
	for _, slugItem := range slugs {
		go func(s string) {
			defer wg.Done()

			parsedURL, err := u.Parse(c.URL)
			if err != nil {
				logger.Log.Error().Err(err).Send()
				return
			}
			parsedURL.Path = s
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

			if enc := string(res.Header.Peek("Content-Encoding")); enc == "gzip" {
				body, err := res.BodyGunzip()
				if err != nil {
					logger.Log.Error().Err(err).Send()
					return
				}

				mu.Lock()
				htmls = append(htmls, body)
				mu.Unlock()
			} else {
				mu.Lock()
				htmls = append(htmls, res.Body())
				mu.Unlock()
			}
		}(slugItem)
	}
	wg.Wait()

	for _, html := range htmls {
		c.digHTML(ctx, html)
	}
}
