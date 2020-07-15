package company

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/SevereCloud/vksdk/api"
	"github.com/gosimple/slug"
	"github.com/nnqq/scr-parser/city"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/mongo"
	"github.com/nnqq/scr-parser/vk"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
	Contacts []vkExecuteContact `json:"contacts"`
	Addr     struct {
		ID      float64 `json:"id"`
		Address string  `json:"address"`
		CityID  float64 `json:"city_id"`
		Title   string  `json:"title"`
	} `json:"addr"`
	City struct {
		ID    float64 `json:"id"`
		Title string  `json:"title"`
	} `json:"city"`
}

type vkExecuteContact struct {
	ID        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	IsClosed  bool    `json:"is_closed"`
	Sex       float64 `json:"sex"`
	Photo200  string  `json:"photo_200"`
}

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func upsertWithRetry(ctx context.Context, doc Company) error {
	opts := options.Update()
	opts.SetUpsert(true)

	for i := 0; i < 10; i += 1 {
		_, err := mongo.Companies.UpdateOne(ctx, bson.M{
			"u": doc.URL,
		}, bson.M{
			"$set": doc,
		}, opts)

		if err == nil {
			break
		}

		if i == 9 {
			logger.Log.Error().Err(err).Send()
			return err
		}

		doc.Slug = strings.Join([]string{
			doc.Slug,
			strconv.Itoa(i + 2),
		}, "-")
	}

	return nil
}

func (c Company) UpdateOrCreate(url, registrar string, registrationDate time.Time) {
	ctx := context.Background()

	uri := strings.Join([]string{
		httpPrefix,
		url,
	}, "")

	doc := Company{
		URL:  uri,
		Slug: slug.Make(url),
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
	if err != nil {
		err = upsertWithRetry(ctx, doc)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
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

	finalURL := resLocation
	if !strings.HasPrefix(resLocation, httpPrefix) || !strings.HasPrefix(resLocation, httpsPrefix) {
		url = strings.Join([]string{
			httpPrefix,
			resLocation,
		}, "")
	}

	doc.Online = true
	doc.URL = finalURL
	doc.Domain.Address = res.RemoteAddr().String()
	doc = digHTML(doc, res.Body())

	err = doc.validate()
	logger.Must(err)

	err = upsertWithRetry(ctx, doc)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}
	logger.Log.Info().
		Bool("online", doc.Online).
		Str("url", doc.URL).
		Msg("website saved")
	return
}

func digHTML(in Company, html []byte) (out Company) {
	out = in

	lowStrHTML := strings.ToLower(string(html))

	foundCity, ok := city.Find(lowStrHTML)
	if ok {
		cityModel := city.City{}
		dbCity, err := cityModel.GetOrCreate(foundCity)
		if err != nil {
			logger.Log.Error().Err(err).Send()
		} else {
			if out.Location == nil {
				out.Location = &location{}
			}
			out.Location.CityID = dbCity.ID
		}
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(lowStrHTML))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	out.Title = strings.TrimSpace(dom.Find("title").Text())
	if len([]rune(out.Title)) > 48 {
		out.Title = string([]rune(out.Title)[:47])
	}

	dom.Find("meta").Each(func(_ int, s *goquery.Selection) {
		prop, _ := s.Attr("name")
		content, _ := s.Attr("content")

		if prop == "description" {
			out.Description = strings.TrimSpace(content)
		}

		// TODO parse og tags
	})

	emailRaw, ok := dom.Find("a[href^='mailto:']").Attr("href")
	if ok {
		out.Email = strings.TrimSpace(strings.Split(emailRaw, "mailto:")[1])
	}

	phoneRaw, ok := dom.Find("a[href^='tel:']").Attr("href")
	if ok {
		phone, err := phoneFromString(phoneRaw)
		if err == nil {
			out.Phone = phone
		}
	}

	if u := getByHrefStart(dom, "http://itunes.apple.com/", "https://itunes.apple.com/",
		"https://www.itunes.apple.com/"); u != "" {
		if out.App == nil {
			out.App = &app{}
		}
		out.App.AppStore = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://play.google.com/", "https://play.google.com/",
		"https://www.play.google.com/"); u != "" {
		if out.App == nil {
			out.App = &app{}
		}
		out.App.GooglePlay = &item{URL: u}
	}

	if u := getByHrefStart(dom, "http://youtube.com/", "https://youtube.com/",
		"https://www.youtube.com/"); u != "" {
		if out.Social == nil {
			out.Social = &social{}
		}
		out.Social.Youtube = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://twitter.com/", "https://twitter.com/",
		"https://www.twitter.com/"); u != "" {
		if out.Social == nil {
			out.Social = &social{}
		}
		out.Social.Twitter = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://facebook.com/", "https://facebook.com/",
		"https://www.facebook.com/"); u != "" {
		if out.Social == nil {
			out.Social = &social{}
		}
		out.Social.Facebook = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://instagram.com/", "https://instagram.com/",
		"https://www.instagram.com/"); u != "" {
		if out.Social == nil {
			out.Social = &social{}
		}
		out.Social.Instagram = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://vk.com/", "https://vk.com/",
		"https://www.vk.com/"); u != "" {
		if out.Social == nil {
			out.Social = &social{}
		}
		out.Social.Vk = &vkItem{URL: u}
	}

	var (
		innFound  = false
		kppFound  = false
		ogrnFound = false
	)
	dom.EachWithBreak(func(_ int, s *goquery.Selection) bool {
		text := s.Text()

		if !innFound {
			index := strings.Index(text, "инн")
			if index != -1 {
				innSubstr := text[index : index+20]
				out.INN, innFound = findInt(innSubstr, "\\s[0-9]{10}\\s")
			}
		}

		if !kppFound {
			index := strings.Index(text, "кпп")
			if index != -1 {
				kppSubstr := text[index : index+18]
				out.KPP, kppFound = findInt(kppSubstr, "\\s[0-9]{9}\\s")
			}
		}

		if !ogrnFound {
			index := strings.Index(text, "огрн")
			if index != -1 {
				ogrnSubstr := text[index : index+26]
				out.OGRN, ogrnFound = findInt(ogrnSubstr, "\\s[0-9]{13}\\s")
			}
		}

		if innFound && kppFound && ogrnFound {
			return false
		}
		return true
	})

	out = digVk(out)

	return
}

func digVk(in Company) (out Company) {
	out = in

	if out.Social != nil && out.Social.Vk != nil && out.Social.Vk.URL != "" {
		execute := vkExecuteRes{}
		groupSlug := strings.TrimSpace(strings.Split(out.Social.Vk.URL, "vk.com/")[1])
		code := fmt.Sprintf(`
			var groups = API.groups.getById({
				group_id: "%s",
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
		`, groupSlug)
		err := vk.UserApi.Execute(code, &execute)
		if err != nil {
			// check is group_id exists, if not - execute allowed to fail
			_, err = vk.UserApi.GroupsGetByID(api.Params{
				"group_ids": groupSlug,
			})
			if err == nil {
				logger.Log.Error().Str("code", code).Msg("execute error")
			}
			return
		}

		if execute.City.Title != "" && execute.City.ID != 0 {
			cityModel := city.City{}
			createdCity, err := cityModel.GetOrCreate(city.NormalCaseCity(execute.City.Title))
			if err != nil {
				logger.Log.Error().Err(err).Send()
			} else {
				if out.Location == nil {
					out.Location = &location{}
				}
				out.Location.CityID = createdCity.ID
			}
		}

		if execute.Addr.Address != "" {
			if out.Location == nil {
				out.Location = &location{}
			}
			out.Location.Address = execute.Addr.Address
		}
		if execute.Addr.Title != "" {
			if out.Location == nil {
				out.Location = &location{}
			}
			out.Location.AddressTitle = execute.Addr.Title
		}

		userMoreFields := map[float64]vkExecuteContact{}
		for _, c := range execute.Contacts {
			userMoreFields[c.ID] = c
		}

		for _, c := range execute.Group.Contacts {
			item := peopleItem{
				VkID:        int(c.UserID),
				Email:       c.Email,
				Description: strings.TrimSpace(c.Desc),
			}

			user, ok := userMoreFields[c.UserID]
			if ok {
				item.FirstName = user.FirstName
				item.LastName = user.LastName
				item.VkIsClosed = user.IsClosed
				item.Sex = int8(user.Sex)
				item.Photo200 = user.Photo200
			}

			phone, err := phoneFromString(c.Phone)
			if err == nil {
				item.Phone = phone
			}

			out.People = append(out.People, &item)
		}

		desc := strings.TrimSpace(execute.Group.Description)

		out.Social.Vk.GroupID = int(execute.Group.ID)
		out.Social.Vk.Name = execute.Group.Name
		out.Social.Vk.ScreenName = execute.Group.ScreenName
		out.Social.Vk.IsClosed = int8(execute.Group.IsClosed)
		out.Social.Vk.Description = desc
		out.Social.Vk.MembersCount = int(execute.Group.MembersCount)
		out.Social.Vk.Photo200 = execute.Group.Photo200

		if execute.Group.Name != "" {
			out.Title = execute.Group.Name
		}
		if execute.Group.Description != "" {
			out.Description = desc
		}
		if execute.Group.Photo200 != "" {
			out.Avatar = execute.Group.Photo200
		}
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

func phoneFromString(in string) (phone int, err error) {
	onlyNumsRx := regexp.MustCompile("[0-9]+")
	numChunks := onlyNumsRx.FindAllString(in, -1)
	if numChunks != nil && len(numChunks) > 0 {
		nums := strings.Join(numChunks, "")
		if string(nums[0]) == "8" {
			nums = strings.Join([]string{"7", nums[1:]}, "")
		}

		return strconv.Atoi(nums)
	}

	err = errors.New("not phone")
	return
}
