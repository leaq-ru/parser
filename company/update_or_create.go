package company

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/logger"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

var client = &fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	ReadTimeout:              10 * time.Second,
	WriteTimeout:             10 * time.Second,
	MaxConnWaitTimeout:       10 * time.Second,
	MaxResponseBodySize:      10 * 1024 * 1024,
}

func (c *Company) UpdateOrCreate(ctx context.Context, url, registrar string, registrationDate time.Time) {
	c.URL = strings.Join([]string{
		httpPrefix,
		url,
	}, "")
	c.Slug = slug.Make(url)
	c.Domain = &domain{
		Registrar:        registrar,
		RegistrationDate: registrationDate,
	}

	mainReq := fasthttp.AcquireRequest()
	mainReq.SetRequestURI(c.URL)
	mainRes := fasthttp.AcquireResponse()
	err := client.DoRedirects(mainReq, mainRes, 3)
	if err != nil {
		err = c.upsertWithRetry(ctx)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		logger.Log.Debug().
			Bool("online", c.Online).
			Str("url", c.URL).
			Msg("website saved")
		return
	}

	c.parseRelatedPages(ctx, client, "contacts", "contact")

	resLocation := c.URL
	if l := string(mainRes.Header.Peek("location")); l != "" {
		resLocation = l
	}
	if l := string(mainRes.Header.Peek("Location")); l != "" {
		resLocation = l
	}

	finalURL := resLocation
	if !strings.HasPrefix(resLocation, httpPrefix) && !strings.HasPrefix(resLocation, httpsPrefix) {
		finalURL = strings.Join([]string{
			httpPrefix,
			resLocation,
		}, "")
	}

	c.Online = true
	c.URL = finalURL
	c.Domain.Address = mainRes.RemoteAddr().String()
	c.digHTML(ctx, mainRes.Body())

	err = c.validate()
	logger.Must(err)

	err = c.upsertWithRetry(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	logger.Log.Debug().
		Bool("online", c.Online).
		Str("url", c.URL).
		Msg("website saved")
	return
}
