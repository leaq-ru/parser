package company

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/nnqq/scr-parser/city"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/rx"
	"strings"
)

func (c *Company) digHTML(ctx context.Context, html []byte) {
	lowStrHTML := strings.ToLower(string(html))

	foundCity, ok := city.Find(lowStrHTML)
	if ok {
		cityModel := city.City{}
		dbCity, err := cityModel.GetOrCreate(ctx, foundCity)
		if err != nil {
			logger.Log.Error().Err(err).Send()
		} else {
			if c.Location == nil {
				c.Location = &location{}
			}
			c.Location.CityID = dbCity.ID
		}
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(lowStrHTML))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	c.Title = capitalize(dom.Find("title").Text())
	if len([]rune(c.Title)) > 48 {
		c.Title = capitalize(string([]rune(c.Title)[:47]))
	}

	dom.Find("meta").Each(func(_ int, s *goquery.Selection) {
		prop, _ := s.Attr("name")
		content, _ := s.Attr("content")

		if prop == "description" {
			c.Description = capitalize(content)
		}
	})

	emailRaw, ok := dom.Find("a[href^='mailto:']").Attr("href")
	if ok {
		c.Email = strings.TrimSpace(strings.Split(emailRaw, "mailto:")[1])
	}
	if c.Email == "" {
		c.Email = strings.TrimSpace(rx.Email.FindString(lowStrHTML))
	}

	phoneRaw, ok := dom.Find("a[href^='tel:']").Attr("href")
	if ok {
		phone, err := rawPhoneToValidPhone(phoneRaw)
		if err == nil {
			c.Phone = phone
		}
	}
	if c.Phone == 0 {
		phone, err := rawPhoneToValidPhone(rx.Phone.FindString(lowStrHTML))
		if err == nil {
			c.Phone = phone
		}
	}

	if u := getByHrefStart(dom, "http://itunes.apple.com/", "https://itunes.apple.com/",
		"https://www.itunes.apple.com/"); u != "" {
		if c.App == nil {
			c.App = &app{}
		}
		c.App.AppStore = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://play.google.com/", "https://play.google.com/",
		"https://www.play.google.com/"); u != "" {
		if c.App == nil {
			c.App = &app{}
		}
		c.App.GooglePlay = &item{URL: u}
	}

	if u := getByHrefStart(dom, "http://youtube.com/", "https://youtube.com/",
		"https://www.youtube.com/"); u != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Youtube = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://twitter.com/", "https://twitter.com/",
		"https://www.twitter.com/"); u != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Twitter = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://facebook.com/", "https://facebook.com/",
		"https://www.facebook.com/"); u != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Facebook = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://instagram.com/", "https://instagram.com/",
		"https://www.instagram.com/"); u != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Instagram = &item{URL: u}
	}
	if u := getByHrefStart(dom, "http://vk.com/", "https://vk.com/",
		"https://www.vk.com/"); u != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Vk = &vkItem{URL: u}
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
				c.INN, innFound = findInt(innSubstr, rx.INN)
			}
		}

		if !kppFound {
			index := strings.Index(text, "кпп")
			if index != -1 {
				kppSubstr := text[index : index+18]
				c.KPP, kppFound = findInt(kppSubstr, rx.KPP)
			}
		}

		if !ogrnFound {
			index := strings.Index(text, "огрн")
			if index != -1 {
				ogrnSubstr := text[index : index+26]
				c.OGRN, ogrnFound = findInt(ogrnSubstr, rx.OGRN)
			}
		}

		if innFound && kppFound && ogrnFound {
			return false
		}
		return true
	})

	c.digVk(ctx)

	return
}
