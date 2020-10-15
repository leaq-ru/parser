package model

import (
	"context"
	"errors"
	userAgent "github.com/EDDYCJY/fake-useragent"
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

func makeSafeFastHTTPClient() *fasthttp.Client {
	return &fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		ReadTimeout:              10 * time.Second,
		WriteTimeout:             10 * time.Second,
		MaxConnWaitTimeout:       10 * time.Second,
		MaxResponseBodySize:      4 * 1024 * 1024,
		ReadBufferSize:           4 * 1024 * 1024,
	}
}

func (c *Company) UpdateOrCreate(ctx context.Context, rawUrl, registrar string, registrationDate time.Time) {
	logger.Log.Debug().Str("rawUrl", rawUrl).Msg("got url to processing")

	url := rawUrl
	if !strings.HasPrefix(url, httpWithSlash) || !strings.HasPrefix(url, httpsWithSlash) {
		url = httpWithSlash + rawUrl
	}

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
		logger.Log.Error().Err(errors.New("invalid url")).Str("url", url).Send()
		return
	}

	c.URL = strings.Join([]string{
		scheme,
		host,
	}, "://")
	c.Slug = slug.Make(host)
	if registrar != "" || registrationDate != (time.Time{}) {
		c.Domain = &domain{
			Registrar:        registrar,
			RegistrationDate: registrationDate,
		}
	}

	mainReq := fasthttp.AcquireRequest()
	mainReq.SetRequestURI(c.URL)
	mainReq.Header.SetUserAgent(userAgent.Random())
	mainRes := fasthttp.AcquireResponse()
	err = makeSafeFastHTTPClient().DoRedirects(mainReq, mainRes, 3)
	if err != nil {
		logger.Log.Debug().
			Err(err).
			Str("url", c.URL).
			Msg("website offline, updated to online=false")

		logger.Err(companySetOffline(ctx, c.Slug))
		return
	}

	c.parseContactsPage(ctx)

	c.Online = true
	c.Domain.Address = mainRes.RemoteAddr().String()

	var body []byte
	if enc := string(mainRes.Header.Peek("Content-Encoding")); enc == "gzip" {
		b, err := mainRes.BodyGunzip()
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		body = b
	} else {
		body = mainRes.Body()
	}

	ogImage := c.digHTML(ctx, body)

	isNoContacts := c.Email == "" && c.Phone == 0
	if isNoContacts || isJunkTitle(c.Title) || isJunkEmail(c.Email) || isJunkPhone(c.Phone) {
		logger.Log.Debug().
			Str("url", c.URL).
			Msg("skip saving junk website")
		return
	}

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

func isJunkEmail(email string) bool {
	switch email {
	case "info@reg.ru",
		"support@beget.com",
		"support@beget.ru",
		"support@mchost.ru",
		"info@timeweb.ru",
		"sales@gobrand.ru",
		"robert@broofa.com":
		return true
	default:
		return false
	}
}

func isJunkTitle(title string) bool {
	if strings.Contains(title, "This website is for sale") ||
		strings.Contains(title, "Ещё один сайт на WordPress") {
		return true
	}

	switch title {
	case "Срок регистрации домена истёк",
		"Срок подключения домена истёк",
		"Продажа облачных доменов для ИТ-проектов.",
		"Домен не прилинкован ни к одной из директорий н":
		return true
	default:
		return false
	}
}

func isJunkPhone(phone int) bool {
	// Продажа облачных доменов для ИТ-проектов.
	return phone == 74503968043
}
