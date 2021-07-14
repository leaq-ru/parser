package company

import (
	"bytes"
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/rx"
	"strings"
	"sync"
)

var space = []byte(" ")

func removeHTMLSpecSymbols(html []byte) []byte {
	noEncodedSpaces := bytes.ReplaceAll(html, []byte("%20"), space)
	return bytes.ReplaceAll(noEncodedSpaces, []byte("&nbsp;"), space)
}

func (c *Company) digHTML(ctx context.Context, rawHTML []byte, setDOMContent, setCity, setCategory bool) (ogImage Link, vkURL string) {
	if len(rawHTML) == 0 {
		err := errors.New("empty HTML")
		logger.Log.Debug().Err(err).Send()
		return
	}

	// <meta property="fb:app_id" content="257953674358265"/>
	html := strings.ReplaceAll(
		string(bytes.ToValidUTF8(
			removeHTMLSpecSymbols(rawHTML),
			space,
		)),
		"257953674358265",
		"")
	if len(html) == 0 {
		return
	}

	if setDOMContent {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return
		}

		c.Title = capitalize(dom.Find("title").Text())
		if len([]rune(c.Title)) > 50 {
			c.Title = capitalize(string([]rune(c.Title)[:50]))
		}

		dom.Find("meta").Each(func(_ int, s *goquery.Selection) {
			name, _ := s.Attr("name")
			property, _ := s.Attr("property")
			content, _ := s.Attr("content")

			if name == "description" {
				desc := capitalize(content)
				if len([]rune(desc)) > 1500 {
					c.Description = string([]rune(desc)[:1500])
				} else {
					c.Description = desc
				}
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
			mailto := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(emailRaw, "mailto:")))
			email := strings.TrimSpace(rx.Email.FindString(mailto))
			if emailSuffixValid(email) {
				c.Email = email
			}
		}
		if c.Email == "" {
			email := strings.TrimSpace(rx.Email.FindString(html))
			if emailSuffixValid(email) {
				c.Email = email
			}
		}

		phoneRaw, ok := dom.Find("a[href^='tel:']").Attr("href")
		if ok {
			phone, err := rawPhoneToValidPhone(phoneRaw)
			if err == nil {
				c.Phone = phone
			}
		}
		if c.Phone == 0 {
			phone, err := rawPhoneToValidPhone(rx.Phone.FindString(html))
			if err == nil {
				c.Phone = phone
			}
		}

		if foundURL := getByHrefStart(dom, "http://apps.apple.com/", "https://apps.apple.com/",
			"https://www.apps.apple.com/"); foundURL != "" {
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
		vkURL = getByHrefStart(dom, "http://vk.com/", "https://vk.com/",
			"https://www.vk.com/")

		var (
			innFound  = c.INN != 0
			kppFound  = c.KPP != 0
			ogrnFound = c.OGRN != 0
		)
		dom.EachWithBreak(func(_ int, s *goquery.Selection) bool {
			text := strings.ToLower(s.Text())

			if !innFound {
				index := strings.Index(text, "инн")
				if index != -1 && len(text) > index+15 {
					innSubstr := text[index : index+15]
					c.INN, innFound = findInt(innSubstr, rx.INN)
				}
			}

			if !kppFound {
				index := strings.Index(text, "кпп")
				if index != -1 && len(text) > index+14 {
					kppSubstr := text[index : index+14]
					c.KPP, kppFound = findInt(kppSubstr, rx.KPP)
				}
			}

			if !ogrnFound {
				index := strings.Index(text, "огрн")
				if index != -1 && len(text) > index+17 {
					ogrnSubstr := text[index : index+17]
					c.OGRN, ogrnFound = findInt(ogrnSubstr, rx.OGRN)
				}
			}

			if innFound && kppFound && ogrnFound {
				return false
			}
			return true
		})
	}

	var wg sync.WaitGroup
	if setCity {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.setCityID(ctx, html)
		}()
	}

	if setCategory {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.setCategoryID(ctx, html)
		}()
	}
	wg.Wait()
	return
}
