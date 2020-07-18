package company

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/logger"
	"github.com/valyala/fasthttp"
	u "net/url"
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
	parsedURL, err := u.Parse(url)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	c.URL = strings.Join([]string{
		parsedURL.Scheme,
		parsedURL.Host,
	}, "://")
	c.Slug = slug.Make(url)
	c.Domain = &domain{
		Registrar:        registrar,
		RegistrationDate: registrationDate,
	}

	mainReq := fasthttp.AcquireRequest()
	mainReq.SetRequestURI(c.URL)
	mainRes := fasthttp.AcquireResponse()
	err = client.DoRedirects(mainReq, mainRes, 3)
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

	c.Online = true
	c.Domain.Address = mainRes.RemoteAddr().String()

	body, err := mainRes.BodyGunzip()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	c.digHTML(ctx, body)

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
