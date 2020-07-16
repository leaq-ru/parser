package company

import (
	"context"
	"github.com/valyala/fasthttp"
	"strings"
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

			req := fasthttp.AcquireRequest()
			req.SetRequestURI(strings.Join([]string{
				strings.TrimSuffix(c.URL, "/"),
				s,
			}, "/"))
			res := fasthttp.AcquireResponse()
			err := client.DoRedirects(req, res, 3)
			if err != nil {
				return
			}

			body, err := res.BodyGunzip()
			if err != nil {
				return
			}

			mu.Lock()
			htmls = append(htmls, body)
			mu.Unlock()
		}(item)
	}
	wg.Wait()

	for _, html := range htmls {
		c.digHTML(ctx, html)
	}
}
