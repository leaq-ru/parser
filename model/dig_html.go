package model

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/rx"
	"strings"
	"unicode/utf8"
)

func (c *Company) digHTML(ctx context.Context, html []byte) (ogImage link) {
	var htmlUTF8 []byte
	if utf8.Valid(html) {
		htmlUTF8 = html
	} else {
		var err error
		htmlUTF8, err = convertToUTF8(html, "windows-1251")
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}
	}

	strHTML := string(htmlUTF8)

	if len(strHTML) == 0 {
		return
	}

	c.setCity(ctx, strHTML)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(strHTML))
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	c.Title = capitalize(dom.Find("title").Text())
	if len([]rune(c.Title)) > 48 {
		c.Title = capitalize(string([]rune(c.Title)[:47]))
	}

	dom.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		property, _ := s.Attr("property")
		content, _ := s.Attr("content")

		if name == "description" {
			c.Description = capitalize(content)
		}
		if property == "og:image" {
			ogImage = toOGImage(strings.TrimSpace(content), c.URL)
		}
	})
	if ogImage == "" {
		imgSrc, ok := dom.Find("img").Attr("src")
		if ok {
			ogImage = toOGImage(imgSrc, c.URL)
		}
	}

	emailRaw, ok := dom.Find("a[href^='mailto:']").Attr("href")
	if ok {
		c.Email = strings.ToLower(strings.TrimSpace(strings.Split(emailRaw, "mailto:")[1]))
	}
	if c.Email == "" {
		c.Email = strings.TrimSpace(rx.Email.FindString(strHTML))
	}

	phoneRaw, ok := dom.Find("a[href^='tel:']").Attr("href")
	if ok {
		phone, err := rawPhoneToValidPhone(phoneRaw)
		if err == nil {
			c.Phone = phone
		}
	}
	if c.Phone == 0 {
		phone, err := rawPhoneToValidPhone(rx.Phone.FindString(strHTML))
		if err == nil {
			c.Phone = phone
		}
	}

	if foundURL := getByHrefStart(dom, "http://itunes.apple.com/", "https://itunes.apple.com/",
		"https://www.itunes.apple.com/"); foundURL != "" {
		if c.App == nil {
			c.App = &app{}
		}
		c.App.AppStore = &item{URL: foundURL}
	}
	if foundURL := getByHrefStart(dom, "http://play.google.com/", "https://play.google.com/",
		"https://www.play.google.com/"); foundURL != "" {
		if c.App == nil {
			c.App = &app{}
		}
		c.App.GooglePlay = &item{URL: foundURL}
	}

	if foundURL := getByHrefStart(dom, "http://youtube.com/", "https://youtube.com/",
		"https://www.youtube.com/"); foundURL != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Youtube = &item{URL: foundURL}
	}
	if foundURL := getByHrefStart(dom, "http://twitter.com/", "https://twitter.com/",
		"https://www.twitter.com/"); foundURL != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Twitter = &item{URL: foundURL}
	}
	if foundURL := getByHrefStart(dom, "http://facebook.com/", "https://facebook.com/",
		"https://www.facebook.com/"); foundURL != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Facebook = &item{URL: foundURL}
	}
	if foundURL := getByHrefStart(dom, "http://instagram.com/", "https://instagram.com/",
		"https://www.instagram.com/"); foundURL != "" {
		if c.Social == nil {
			c.Social = &social{}
		}
		c.Social.Instagram = &item{URL: foundURL}
	}
	if foundURL := getByHrefStart(dom, "http://vk.com/", "https://vk.com/",
		"https://www.vk.com/"); foundURL != "" {
		c.digVk(ctx, foundURL)
	}

	var (
		innFound  = false
		kppFound  = false
		ogrnFound = false
	)
	dom.EachWithBreak(func(_ int, s *goquery.Selection) bool {
		text := strings.ToLower(s.Text())

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

	return
}
