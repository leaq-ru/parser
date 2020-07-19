package company

import (
	"context"
	"github.com/nnqq/scr-parser/logger"
	"github.com/valyala/fasthttp"
	u "net/url"
	"sync"
)

func (c *Company) parseRelatedPages(ctx context.Context, client *fasthttp.Client, slugs ...string) {
	var (
		mu    sync.Mutex
		htmls [][]byte
	)

	wg := sync.WaitGroup{}
	wg.Add(len(slugs))
	for _, item := range slugs {
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
			res := fasthttp.AcquireResponse()
			err = client.DoRedirects(req, res, 3)
			if err != nil {
				logger.Log.Error().Err(err).Send()
				return
			}

			mu.Lock()
			htmls = append(htmls, res.Body())
			mu.Unlock()
		}(item)
	}
	wg.Wait()

	for _, html := range htmls {
		c.digHTML(ctx, html)
	}
}
