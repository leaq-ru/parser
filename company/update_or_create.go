package company

import (
	"context"
	"errors"
	userAgent "github.com/EDDYCJY/fake-useragent"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/call"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/post"
	"github.com/nnqq/scr-proto/codegen/go/image"
	"github.com/nnqq/scr-proto/codegen/go/technology"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		MaxResponseBodySize:      4 * 1024 * 1024,
		ReadBufferSize:           4 * 1024 * 1024,
	}
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

	lowRawURL := strings.ToLower(rawURL)

	url := lowRawURL
	if !strings.HasPrefix(url, httpWithSlash) || !strings.HasPrefix(url, httpsWithSlash) {
		url = httpWithSlash + lowRawURL
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
	if host == "" || host == "leaq.ru" {
		logger.Log.Error().Err(errors.New("invalid url")).Str("url", url).Send()
		return
	}

	// to process .рф sites
	maybePunycodeURL := strings.Join([]string{
		scheme,
		host,
	}, "://")

	unicodeHost, err := idna.New().ToUnicode(host)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	c.Slug = slug.Make(unicodeHost)

	if registrar != "" || registrationDate != (time.Time{}) {
		c.Domain = &domain{
			Registrar:        registrar,
			RegistrationDate: registrationDate,
		}
	}

	mainReq := fasthttp.AcquireRequest()
	mainReq.SetRequestURI(maybePunycodeURL)
	mainReq.Header.SetUserAgent(userAgent.Random())
	mainRes := fasthttp.AcquireResponse()
	pageSpeedStart := time.Now()
	err = makeSafeFastHTTPClient().DoRedirects(mainReq, mainRes, 3)
	pageSpeed := time.Since(pageSpeedStart).Milliseconds()
	if err != nil {
		logger.Log.Debug().
			Err(err).
			Str("url", maybePunycodeURL).
			Msg("website offline, updated to online=false")

		logger.Err(companySetOffline(ctx, c.Slug))
		return
	}

	c.parseContactsPage(ctx, maybePunycodeURL)

	// made request with punycode, now set to human readable url
	c.URL = strings.Join([]string{
		scheme,
		unicodeHost,
	}, "://")
	c.Online = true
	c.PageSpeed = uint32(pageSpeed)
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

	c.digVk(ctx, vkURL)
	isNoVkGroup := c.GetSocial().GetVk().GetGroupId() == 0
	if isNoContacts && isNoVkGroup {
		logger.Log.Debug().
			Str("url", c.URL).
			Msg("skip saving junk website, no contacts and vk group found")
		return
	}

	c.digHTML(ctx, body, false, true, true)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var oldComp Company
		errFindOne := mongo.Companies.FindOne(ctx, bson.M{
			"u": c.URL,
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

	var techOIDs []primitive.ObjectID
	go func(url string) {
		defer wg.Done()
		techs, errFind := call.Technology.Find(ctx, &technology.FindRequest{Url: url})
		if errFind != nil {
			logger.Log.Error().Err(errFind).Send()
			return
		}

		for _, id := range techs.GetIds() {
			oID, errOID := primitive.ObjectIDFromHex(id)
			if errOID != nil {
				logger.Log.Error().Err(errOID).Send()
				continue
			}
			techOIDs = append(techOIDs, oID)
		}
	}(c.URL)
	wg.Wait()

	c.TechnologyIDs = techOIDs

	err = c.upsertWithRetry(ctx)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	logger.Log.Debug().
		Bool("online", c.Online).
		Str("url", c.URL).
		Msg("website saved")

	if c.Social != nil && c.Social.Vk != nil {
		var comp Company
		err = mongo.Companies.FindOne(ctx, Company{
			URL: c.URL,
		}, options.FindOne().SetProjection(bson.M{
			"_id": 1,
		})).Decode(&comp)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		startReplace := time.Now()
		err = post.ReplaceMany(ctx, comp.ID, c.Social.Vk.GroupID)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
		logger.Log.Debug().
			Dur("ms", time.Since(startReplace)).
			Msg("company posts replaced with new one")
	}
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
		"robert@broofa.com":
		return true
	default:
		return false
	}
}

func isJunkTitle(title string) bool {
	t := strings.ToLower(title)

	if strings.Contains(t, "this website is for sale") ||
		strings.Contains(t, "ещё один сайт на wordpress") {
		return true
	}

	switch t {
	case "срок регистрации домена истёк",
		"срок подключения домена истёк",
		"продажа облачных доменов для ит-проектов.",
		"домен не прилинкован ни к одной из директорий н":
		return true
	default:
		return false
	}
}

func isJunkPhone(phone int) bool {
	// Продажа облачных доменов для ИТ-проектов.
	return phone == 74503968043
}
