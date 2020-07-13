package model

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	validation "github.com/go-ozzo/ozzo-validation"
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
type link = string

type Company struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"u,omitempty"`
	Title       string             `bson:"t,omitempty"`
	Type        string             `bson:"ty,omitempty"`
	Email       string             `bson:"e,omitempty"`
	Description string             `bson:"d,omitempty"`
	Online      bool               `bson:"o,omitempty"`
	Phone       int                `bson:"p,omitempty"`
	INN         int                `bson:"i,omitempty"`
	KPP         int                `bson:"k,omitempty"`
	OGRN        int                `bson:"og,omitempty"`
	Domain      *domain            `bson:"do,omitempty"`
	Avatar      link               `bson:"a,omitempty"`
	Location    *location          `bson:"l,omitempty"`
	App         *app               `bson:"ap,omitempty"`
	Social      *social            `bson:"s,omitempty"`
	People      []*peopleItem      `bson:"pe,omitempty"`
}

type peopleItem struct {
	VkID       int    `bson:"v,omitempty"`
	FirstName  string `bson:"f,omitempty"`
	LastName   string `bson:"l,omitempty"`
	VkIsClosed bool   `bson:"vc,omitempty"`
	Sex        int8   `bson:"s,omitempty"`
	Photo200   link   `bson:"ph,omitempty"`
	Phone      int    `bson:"p,omitempty"`
	Email      string `bson:"e,omitempty"`
}

// move to another collection
type location struct {
	VkCityID     int    `bson:"v,omitempty"`
	Address      string `bson:"a,omitempty"`
	AddressTitle string `bson:"at,omitempty"`
	CityTitle    string `bson:"c,omitempty"`
}

type domain struct {
	Address          string    `bson:"a,omitempty"`
	Registrar        string    `bson:"r,omitempty"`
	RegistrationDate time.Time `bson:"rd,omitempty"`
}

type social struct {
	Vk        *vkItem `bson:"v,omitempty"`
	Instagram *item   `bson:"i,omitempty"`
	Twitter   *item   `bson:"t,omitempty"`
	Youtube   *item   `bson:"y,omitempty"`
	Facebook  *item   `bson:"f,omitempty"`
}

type vkItem struct {
	URL          string `bson:"u,omitempty"`
	GroupID      int    `bson:"g,omitempty"`
	Name         string `bson:"n,omitempty"`
	ScreenName   string `bson:"s,omitempty"`
	IsClosed     int8   `bson:"i,omitempty"`
	Description  string `bson:"d,omitempty"`
	MembersCount int    `bson:"m,omitempty"`
	Photo200     link   `bson:"p,omitempty"`
}

type app struct {
	AppStore   *item `bson:"a,omitempty"`
	GooglePlay *item `bson:"g,omitempty"`
}

type item struct {
	URL string `bson:"u,omitempty"`
}

//{
//	"response": {
//		"group": {
//			"id": 144090016,
//			"name": "Каркасные авточехлы dress4car | +7 904 0555 202",
//			"screen_name": "dress4car",
//			"is_closed": 0,
//			"type": "page",
//			"is_admin": 1,
//			"admin_level": 3,
//			"is_member": 0,
//			"is_advertiser": 1,
//			"addresses": {
//				"is_enabled": true,
//				"main_address_id": 1784
//			},
//			"description": "Пошив и установка авточехлов из экокожи...",
//			"members_count": 37026,
//			"contacts": [{
//				"user_id": 421825761,
//				"desc": "Консультация и заказ",
//				"phone": "+7 904 0555 202"
//			}],
//			"photo_50": "https://sun1-25.u...NyrDcrl2Q.jpg?ava=1",
//			"photo_100": "https://sun1-23.u...do0zcQzLo.jpg?ava=1",
//			"photo_200": "https://sun1-18.u...CcGj_8RgM.jpg?ava=1"
//		},
//		"contacts": [{
//			"id": 421825761,
//			"first_name": "Андрей",
//			"last_name": "Аверьянов",
//			"is_closed": false,
//			"can_access_closed": true,
//			"sex": 2,
//			"photo_200": "https://sun1-83.u...BLFe0d6k4.jpg?ava=1"
//		}],
//		"addr": {
//			"id": 1784,
//			"address": "ул.Дачная, 1-А",
//			"city_id": 95,
//			"title": "Детейлинг центр AutoDOL"
//		},
//		"city": {
//			"id": 95,
//			"title": "Нижний Новгород"
//		}
//	}
//}
type vkExecuteRes struct {
	Response struct {
		Group struct {
			ID           float64 `json:"id"`
			Name         string  `json:"name"`
			ScreenName   string  `json:"screen_name"`
			IsClosed     float64 `json:"is_closed"`
			Description  string  `json:"description"`
			MembersCount float64 `json:"members_count"`
			Contacts     []struct {
				UserID float64 `json:"user_id"`
				Desc   string  `json:"desc"`
				Phone  string  `json:"phone"`
				Email  string  `json:"email"`
			} `json:"contacts"`
			Photo200 string `json:"photo_200"`
		} `json:"group"`
		Contacts []struct {
			ID        float64 `json:"id"`
			FirstName string  `json:"first_name"`
			LastName  string  `json:"last_name"`
			IsClosed  bool    `json:"is_closed"`
			Sex       float64 `json:"sex"`
			Photo200  string  `json:"photo_200"`
		} `json:"contacts"`
		Addr struct {
			ID      float64 `json:"id"`
			Address string  `json:"address"`
			CityID  float64 `json:"city_id"`
			Title   string  `json:"title"`
		} `json:"addr"`
		City struct {
			ID    float64 `json:"id"`
			Title string  `json:"title"`
		} `json:"city"`
	} `json:"response"`
	ExecuteErrors []struct {
		Method    string  `json:"method"`
		ErrorCode float64 `json:"error_code"`
		ErrorMsg  string  `json:"error_msg"`
	}
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

func (c Company) Upsert(url, registrar string, registrationDate time.Time) {
	uri := strings.Join([]string{
		"http://",
		url,
	}, "")

	doc := Company{
		URL: uri,
		Domain: &domain{
			Registrar:        registrar,
			RegistrationDate: registrationDate,
		},
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	res := fasthttp.AcquireResponse()

	sec3 := 3 * time.Second
	client := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		ReadTimeout:              sec3,
		WriteTimeout:             sec3,
		MaxConnWaitTimeout:       sec3,
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

	if u := getByHrefStart(doc, "http://itunes.apple.com/", "https://itunes.apple.com/",
		"https://www.itunes.apple.com/"); u != "" {
		if res.App == nil {
			res.App = &app{}
		}
		res.App.AppStore = &item{URL: u}
	}
	if u := getByHrefStart(doc, "http://play.google.com/", "https://play.google.com/",
		"https://www.play.google.com/"); u != "" {
		if res.App == nil {
			res.App = &app{}
		}
		res.App.GooglePlay = &item{URL: u}
	}

	if u := getByHrefStart(doc, "http://youtube.com/", "https://youtube.com/",
		"https://www.youtube.com/"); u != "" {
		if res.Social == nil {
			res.Social = &social{}
		}
		res.Social.Youtube = &item{URL: u}
	}
	if u := getByHrefStart(doc, "http://twitter.com/", "https://twitter.com/",
		"https://www.twitter.com/"); u != "" {
		if res.Social == nil {
			res.Social = &social{}
		}
		res.Social.Twitter = &item{URL: u}
	}
	if u := getByHrefStart(doc, "http://facebook.com/", "https://facebook.com/",
		"https://www.facebook.com/"); u != "" {
		if res.Social == nil {
			res.Social = &social{}
		}
		res.Social.Facebook = &item{URL: u}
	}
	if u := getByHrefStart(doc, "http://instagram.com/", "https://instagram.com/",
		"https://www.instagram.com/"); u != "" {
		if res.Social == nil {
			res.Social = &social{}
		}
		res.Social.Instagram = &item{URL: u}
	}
	if u := getByHrefStart(doc, "http://vk.com/", "https://vk.com/",
		"https://www.vk.com/"); u != "" {
		if res.Social == nil {
			res.Social = &social{}
		}
		res.Social.Vk = &vkItem{URL: u}
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

	if res.Social != nil && res.Social.Vk != nil && res.Social.Vk.URL != "" {
		execRes := vkExecuteRes{}
		err := vk.Api.Execute(fmt.Sprintf(`
			var groups = API.groups.getById({
				group_id: %s,
				fields: "addresses,description,members_count,contacts",
				v: "5.120",
			});
			var group = groups[0];

			var contacts = API.users.get({
				user_ids: group.contacts@.user_id,
				fields: "photo_200,sex",
				v: "5.120",
			});

			var addrs = API.groups.getAddresses({
				group_id: group.id,
				address_ids: group.addresses.main_address_id,
				fields: "title,address,city_id",
				count: 1,
				v: "5.120",
			});
			var addr = addrs.items[0];

			var cities = API.database.getCitiesById({
				city_ids: addr.city_id,
				v: "5.120",
			});
			var city = cities[0];

			return {
				group: group,
				contacts: contacts,
				addr: addr,
				city: city,
			};
		`, strings.Split(res.Social.Vk.URL, "vk.com/")[1]), &execRes)
		if err != nil {
			logger.Log.Error().Stack().Err(err).Send()
			return
		}
		if len(execRes.ExecuteErrors) != 0 {
			logger.Log.Error().Stack().Msgf("%+v\n", execRes.ExecuteErrors)
			return
		}

		// TODO parse execRes data
	}

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
