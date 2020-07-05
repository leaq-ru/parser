package model

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/SevereCloud/vksdk/api"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/nnqq/scr-parser/config"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/vk"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Direct link .jpg
type avatar = string

type Company struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	PeopleIDs []primitive.ObjectID `bson:"pi,omitempty"`
	URL       string               `bson:"u,omitempty"`
	Type      string               `bson:"t,omitempty"`
	Email     string               `bson:"e,omitempty"`
	Online    bool                 `bson:"o,omitempty"`
	Phone     int                  `bson:"p,omitempty"`
	INN       int                  `bson:"i,omitempty"`
	KPP       int                  `bson:"k,omitempty"`
	OGRN      int                  `bson:"og,omitempty"`
	Domain    domain               `bson:"d,omitempty"`
	Avatar    avatar               `bson:"a,omitempty"`
	Location  location             `bson:"l,omitempty"`
	Apps      apps                 `bson:"ap,omitempty"`
	Social    social               `bson:"s,omitempty"`
}

type location struct {
	Address string             `bson:"a,omitempty"`
	CityID  primitive.ObjectID `bson:"c,omitempty"`
}

type domain struct {
	Address          string    `bson:"a,omitempty"`
	Registrar        string    `bson:"r,omitempty"`
	RegistrationDate time.Time `bson:"rd,omitempty"`
}

type social struct {
	Vk        item `bson:"v,omitempty"`
	Instagram item `bson:"i,omitempty"`
	Twitter   item `bson:"t,omitempty"`
	Youtube   item `bson:"y,omitempty"`
	Facebook  item `bson:"f,omitempty"`
}

type apps struct {
	AppStore   item `bson:"a,omitempty"`
	GooglePlay item `bson:"g,omitempty"`
}

type item struct {
	URL string `bson:"u,omitempty"`
}

type vkExecuteRes []struct {
	UserId int    `json:"user_id"`
	Avatar string `json:"avatar"`
	Desc   string `json:"desc"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

func (c Company) validate() error {
	err := validation.ValidateStruct(
		&c,
		validation.Field(&c.ID, validation.Required),
		validation.Field(&c.URL, validation.Required),
		validation.Field(&c.Online, validation.Required),
	)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(
		&c.Domain,
		validation.Field(&c.Domain.Address, validation.Required),
		validation.Field(&c.Domain.Registrar, validation.Required),
		validation.Field(&c.Domain.RegistrationDate, validation.Required),
	)
}

func (c Company) Create(url, registrar string, registrationDate time.Time) {
	uri := strings.Join([]string{
		"http://",
		url,
	}, "")

	doc := Company{
		URL: uri,
		Domain: domain{
			Registrar:        registrar,
			RegistrationDate: registrationDate,
		},
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	res := fasthttp.AcquireResponse()

	client := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		ReadTimeout:              time.Second,
		WriteTimeout:             time.Second,
		MaxConnWaitTimeout:       time.Second,
		MaxResponseBodySize:      math.MaxInt64,
	}
	err := client.DoRedirects(req, res, 3)
	opts := options.UpdateOptions{}
	opts.SetUpsert(true)
	if err != nil {
		_, err := mongo.Companies.UpdateOne(context.Background(), bson.M{
			"u": doc.URL,
		}, bson.M{
			"$set": doc,
		}, &opts)
		if err != nil {
			logger.Log.Error().Err(err).Send()
		}

		logger.Log.Info().Err(err).
			Bool("online", doc.Online).
			Str("url", doc.URL).
			Msg("website saved")
		return
	}

	resLocation := uri
	if l := string(res.Header.Peek("location")); l != "" {
		resLocation = l
	}
	if l := string(res.Header.Peek("Location")); l != "" {
		resLocation = l
	}

	doc.Online = true
	doc.URL = resLocation
	doc.Domain.Address = res.RemoteAddr().String()
	doc = digHTML(doc, res.Body())

	err = doc.validate()
	logger.Must(err)

	_, err = mongo.Companies.UpdateOne(context.Background(), bson.M{
		"u": doc.URL,
	}, bson.M{
		"$set": doc,
	}, &opts)
	if err != nil {
		logger.Log.Error().Err(err).Send()
	}
	logger.Log.Info().
		Bool("online", doc.Online).
		Str("url", doc.URL).
		Msg("website saved")
	return
}

func digHTML(in Company, html []byte) (res Company) {
	res = in

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return
	}

	emailRaw, ok := doc.Find("a[href^='mailto:']").Attr("href")
	if ok {
		res.Email = strings.TrimSpace(strings.Split(emailRaw, "mailto:")[1])
	}

	phoneRaw, ok := doc.Find("a[href^='tel:']").Attr("href")
	if ok {
		onlyNumsRx := regexp.MustCompile("[0-9]+")
		numChunks := onlyNumsRx.FindAllString(strings.Split(phoneRaw, "tel:")[1], -1)
		if numChunks != nil && len(numChunks) > 0 {
			nums := strings.Join(numChunks, "")
			if string(nums[0]) == "8" {
				nums = strings.Join([]string{"7", nums[1:]}, "")
			}

			p, err := strconv.Atoi(nums)
			if err == nil {
				res.Phone = p
			}
		}
	}

	res.Apps.AppStore.URL = getByHrefStart(doc, "http://itunes.apple.com/", "https://itunes.apple.com/",
		"https://www.itunes.apple.com/")
	res.Apps.GooglePlay.URL = getByHrefStart(doc, "http://play.google.com/", "https://play.google.com/",
		"https://www.play.google.com/")

	res.Social.Youtube.URL = getByHrefStart(doc, "http://youtube.com/", "https://youtube.com/",
		"https://www.youtube.com/")
	res.Social.Twitter.URL = getByHrefStart(doc, "http://twitter.com/", "https://twitter.com/",
		"https://www.twitter.com/")
	res.Social.Facebook.URL = getByHrefStart(doc, "http://facebook.com/", "https://facebook.com/",
		"https://www.facebook.com/")
	res.Social.Instagram.URL = getByHrefStart(doc, "http://instagram.com/", "https://instagram.com/",
		"https://www.instagram.com/")
	res.Social.Vk.URL = getByHrefStart(doc, "http://vk.com/", "https://vk.com/",
		"https://www.vk.com/")

	if res.Social.Vk.URL != "" {
		execRes := vkExecuteRes{}

		err := vk.Api.Execute(fmt.Sprintf(`
			var group = API.groups.getById({
				"group_id": %s,
				"fields": "description,place,members_count,contacts",
				"v": "5.120",
			});

			var users = API.users.get({
				"user_ids": group[0].contacts@.user_id,
				"fields": "city,photo_200,sex",
				"v": "5.120",
			});

			return {
				"group": group,
				"users": users,
			};
		`, strings.Split(res.Social.Vk.URL, "vk.com/")[1]), &execRes)
	}

	var (
		innFound  = false
		kppFound  = false
		ogrnFound = false
	)
	doc.EachWithBreak(func(_ int, s *goquery.Selection) bool {
		text := strings.ToLower(s.Text())

		if !innFound {
			index := strings.Index(text, "инн")
			if index != -1 {
				innSubstr := text[index : index+20]
				res.INN, innFound = findInt(innSubstr, "\\s[0-9]{10}\\s")
			}
		}

		if !kppFound {
			index := strings.Index(text, "кпп")
			if index != -1 {
				kppSubstr := text[index : index+18]
				res.KPP, kppFound = findInt(kppSubstr, "\\s[0-9]{9}\\s")
			}
		}

		if !ogrnFound {
			index := strings.Index(text, "огрн")
			if index != -1 {
				ogrnSubstr := text[index : index+26]
				res.OGRN, ogrnFound = findInt(ogrnSubstr, "\\s[0-9]{13}\\s")
			}
		}

		if innFound && kppFound && ogrnFound {
			return false
		}
		return true
	})

	return
}

func findInt(text string, pattern string) (result int, found bool) {
	rx := regexp.MustCompile(pattern)
	noSpaces := regexp.MustCompile("\\s")
	r, err := strconv.Atoi(noSpaces.ReplaceAllString(rx.FindString(text), ""))
	if err == nil {
		result = r
		found = true
	}
	return
}

func getByHrefStart(doc *goquery.Document, starts ...string) (hrefAttr string) {
	for _, elem := range starts {
		h, ok := doc.Find(fmt.Sprintf("a[href^='%s']", elem)).Attr("href")
		if ok && h != elem {
			hrefAttr = h
			return
		}
	}
	return
}
