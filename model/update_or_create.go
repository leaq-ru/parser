package model

import (
	"context"
	"errors"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
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

	scheme := parsedURL.Scheme
	if scheme == "" {
		scheme = http
	}
	host := parsedURL.Host
	if host == "" {
		if parsedURL.Path == "" {
			logger.Log.Error().Err(errors.New("invalid url")).Str("url", url).Send()
			return
		}
		host = parsedURL.Path
	}

	c.URL = strings.Join([]string{
		scheme,
		host,
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

	ogImage := c.digHTML(ctx, mainRes.Body())

	oldComp := Company{}
	err = mongo.Companies.FindOne(ctx, bson.M{
		"u": c.URL,
	}).Decode(&oldComp)
	if err != nil {
		if errors.Is(err, m.ErrNoDocuments) {
			if ogImage != "" {
				err = c.setAvatar(ctx, ogImage)
				if err != nil {
					logger.Log.Debug().Str("ogImage", string(ogImage)).Err(err).Send()
				}
			}
		} else {
			logger.Log.Error().Err(err).Send()
			return
		}
	} else {
		if ogImage != "" {
			// try to set new avatar, if OK, then delete old from S3
			err = c.setAvatar(ctx, ogImage)
			if err != nil {
				logger.Log.Debug().Err(err).Send()
			} else {
				if oldComp.Avatar != "" {
					_, err = call.Image.Remove(ctx, &image.RemoveRequest{
						S3Url: string(oldComp.Avatar),
					})
					if err != nil {
						logger.Log.Error().Err(err).Send()
						return
					}
				}
			}
		}
	}

	err = c.validate()
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

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
