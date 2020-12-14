package company

import (
	"context"
	"errors"
	userAgent "github.com/EDDYCJY/fake-useragent"
	"github.com/nnqq/scr-parser/logger"
	"github.com/valyala/fasthttp"
	u "net/url"
	"strings"
	"sync"
)

var ErrBlindHintNotOK = errors.New("got not 200 code on blind slug hit")

func parseSlug(url, slug string) (html []byte, err error) {
	parsedURL, err := u.Parse(url)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	parsedURL.Path = slug
	withSlug := parsedURL.String()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(withSlug)
	req.Header.SetUserAgent(userAgent.Random())
	defer fasthttp.ReleaseRequest(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	err = makeSafeFastHTTPClient().DoRedirects(req, res, 3)
	if err != nil {
		if !errors.Is(err, fasthttp.ErrTooManyRedirects) {
			logger.Log.Error().Err(err).Send()
		}
		return
	}

	if res.StatusCode() != fasthttp.StatusOK {
		err = ErrBlindHintNotOK
		logger.Log.Debug().Str("withSlug", withSlug).Err(err).Send()
		return
	}
	logger.Log.Debug().Str("withSlug", withSlug).Msg("blind related hit")

	html = res.Body()
	return
}

func (c *Company) parseContactsPage(ctx context.Context, url string) {
	slugs := []string{"contacts", "kontakty", "contact-us", "contact"}

	// .рф
	if strings.HasSuffix(url, ".xn--p1ai") {
		// контакты
		slugs = append(slugs, "xn--80atbkezc6e")
	}

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		htmls [][]byte
	)
	wg.Add(len(slugs))
	for _, _slug := range slugs {
		go func(slug string) {
			defer wg.Done()

			html, err := parseSlug(url, slug)
			if err != nil {
				if !errors.Is(err, ErrBlindHintNotOK) {
					logger.Log.Error().Err(err).Send()
				}
				return
			}

			mu.Lock()
			htmls = append(htmls, html)
			mu.Unlock()
		}(_slug)
	}
	wg.Wait()

	for _, html := range htmls {
		c.digHTML(ctx, html, true, true, false)
	}
}
