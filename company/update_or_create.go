package company

import (
	"bytes"
	"context"
	"errors"
	userAgent "github.com/EDDYCJY/fake-useragent"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/post"
	"github.com/nnqq/scr-parser/ptr"
	"github.com/nnqq/scr-parser/stan"
	"github.com/nnqq/scr-proto/codegen/go/event"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/valyala/fasthttp"
	m "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/idna"
	u "net/url"
	"strings"
	"sync"
	"time"
)

func makeSafeFastHTTPClient() *fasthttp.Client {
	return &fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		ReadTimeout:              5 * time.Second,
		WriteTimeout:             5 * time.Second,
		MaxConnWaitTimeout:       5 * time.Second,
		MaxResponseBodySize:      2 * 1024 * 1024,
		ReadBufferSize:           2 * 1024 * 1024,
	}
}

func makeURL(host string) string {
	return strings.Join([]string{
		http,
		host,
	}, "://")
}

func (c *Company) UpdateOrCreate(ctx context.Context, rawURL, registrar string, registrationDate time.Time) {
	start := time.Now()
	logger.Log.Debug().
		Str("rawURL", rawURL).
		Msg("url processing start")
	defer func() {
		logger.Log.Debug().
			Str("rawURL", rawURL).
			Dur("ms", time.Since(start)).
			Msg("url processing finish")
	}()

	parsedURL, err := u.Parse(makeURL(strings.ToLower(rawURL)))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	host := parsedURL.Host
	if host == "" || host == "leaq.ru" {
		logger.Log.Error().Err(errors.New("invalid url")).Str("rawURL", rawURL).Send()
		return
	}

	// to process .рф sites
	punycodeURL := makeURL(host)

	if registrar != "" {
		if c.Domain == nil {
			c.Domain = &domain{}
		}
		c.Domain.Registrar = registrar
	}
	if !registrationDate.IsZero() {
		if c.Domain == nil {
			c.Domain = &domain{}
		}
		c.Domain.RegistrationDate = registrationDate
	}

	mainReq := fasthttp.AcquireRequest()
	mainReq.SetRequestURI(punycodeURL)
	mainReq.Header.SetUserAgent(userAgent.Random())
	defer fasthttp.ReleaseRequest(mainReq)

	mainRes := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(mainRes)

	pageSpeedStart := time.Now()
	err = makeSafeFastHTTPClient().DoRedirects(mainReq, mainRes, 3)
	pageSpeed := time.Since(pageSpeedStart).Milliseconds()
	if err != nil {
		logger.Log.Debug().
			Err(err).
			Str("url", punycodeURL).
			Msg("website offline, updated to online=false")

		logger.Err(companySetOffline(ctx, punycodeURL))
		return
	}

	realPunycodeHost := string(bytes.TrimSuffix(bytes.TrimPrefix(mainReq.URI().Host(), []byte("www.")), []byte(":443")))
	realPunycodeURL := makeURL(realPunycodeHost)

	realUnicodeHost, err := idna.New().ToUnicode(realPunycodeHost)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	realUnicodeURL := makeURL(realUnicodeHost)

	err = mongo.Companies.FindOne(ctx, Company{
		URL:      realUnicodeURL,
		Verified: true,
	}).Err()
	if err == nil {
		logger.Log.Debug().Str("url", realUnicodeURL).Msg("company has human owner, skip reindex")
		return
	}
	if !errors.Is(err, m.ErrNoDocuments) {
		logger.Log.Error().Err(err).Send()
		return
	}

	c.parseContactsPage(ctx, realPunycodeURL)

	c.Slug = slug.Make(realUnicodeHost)

	// made requests with punycode, now set to human readable url
	c.URL = realUnicodeURL
	c.Online = ptr.Bool(true)
	c.PageSpeed = uint32(pageSpeed)
	if c.Domain == nil {
		c.Domain = &domain{}
	}
	c.Domain.Address = mainRes.RemoteAddr().String()

	body := mainRes.Body()

	ogImage, vkURL := c.digHTML(ctx, body, true, false, false)

	isNoContacts := c.Email == "" && c.Phone == 0
	isNoVkURL := vkURL == ""
	isNoData := isNoContacts && isNoVkURL
	if isNoData ||
		isJunkTitle(c.Title) ||
		isJunkDescription(c.Description) ||
		isJunkEmail(c.Email) ||
		isJunkPhone(c.Phone) {
		logger.Log.Debug().
			Str("url", c.URL).
			Msg("skip saving junk website")
		return
	}

	c.DigVk(ctx, vkURL)
	isNoVkGroup := c.GetSocial().GetVk().GetGroupId() == 0
	if isNoContacts && isNoVkGroup {
		logger.Log.Debug().
			Str("url", c.URL).
			Msg("skip saving junk website, no contacts and vk group found")
		return
	}

	c.digHTML(ctx, body, false, true, true)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		var oldComp Company
		errFindOne := mongo.Companies.FindOne(ctx, Company{
			URL: c.URL,
		}).Decode(&oldComp)

		if errFindOne != nil {
			if errors.Is(errFindOne, m.ErrNoDocuments) {
				if ogImage != "" {
					err = c.setAvatar(ctx, ogImage)
					if err != nil {
						logger.Log.Debug().Str("ogImage", string(ogImage)).Err(err).Send()
					}
				}
			} else {
				logger.Log.Error().Err(errFindOne).Send()
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
	}()

	c.withDNS(ctx)

	err = c.upsertWithRetry(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	id, err := getIDByURL(ctx, c.URL)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	if c.Social != nil && c.Social.Vk != nil {
		startReplace := time.Now()
		err = post.ReplaceMany(ctx, id, c.Social.Vk.GroupID, false)
		if err != nil {
			logger.Log.Error().Err(err).Send()
		} else {
			logger.Log.Debug().
				Dur("ms", time.Since(startReplace)).
				Msg("company posts replaced with new one")
		}
	}

	err = stan.ProduceCompanyNew(&event.CompanyNew{
		CompanyId: id.Hex(),
		Url:       c.URL,
	})
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	logger.Log.Debug().
		Str("url", c.URL).
		Msg("website saved")
	return
}

func isJunkDescription(desc string) bool {
	return strings.Contains(strings.ToLower(desc), "domain is parked by service domainparking.ru")
}

func isJunkEmail(email string) bool {
	switch email {
	case "info@reg.ru",
		"support@beget.com",
		"support@beget.ru",
		"support@mchost.ru",
		"all@pwhost.ru",
		"sales@firstvds.ru",
		"support@hostline.ru",
		"support@axelname.ru",
		"info@timeweb.ru",
		"billing@hostia.ru",
		"sales@ispserver.com",
		"sales@gobrand.ru",
		"robert@broofa.com",
		"donate@opencart.com",
		"support@sim-networks.com":
		return true
	default:
		return false
	}
}

func stringContains(str string, substrs ...string) bool {
	for _, substr := range substrs {
		if strings.Contains(str, substr) {
			return true
		}
	}
	return false
}

func isJunkTitle(title string) bool {
	return stringContains(strings.ToLower(title),
		"this website is for sale",
		"ещё один сайт на wordpress",
		"продается домен",
		"домен продается",
		"продам домен",
		"доменное имя временно заблокировано",
		"срок регистрации домена истёк",
		"срок подключения домена истёк",
		"срок регистрации домена закончился",
		"продажа облачных доменов для ит-проектов",
		"домен не прилинкован ни к одной из директорий",
		"не опубликован",
		"spaceweb",
		"припаркован",
		"ошибка 403",
		"your store",
		"сайт создан",
		"сайт заблокирован",
		"сайт временно заблокирован",
		"access forbidden",
		"страница не найдена",
		"проститут",
		"шлюх",
		"доставка алкоголя")
}

func isJunkPhone(phone int) bool {
	// Продажа облачных доменов для ИТ-проектов.
	return phone == 74503968043
}
