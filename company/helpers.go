package company

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/rx"
	"golang.org/x/net/idna"
	u "net/url"
	"regexp"
	"strconv"
	"strings"
)

func capitalize(in string) string {
	text := strings.TrimSpace(in)
	if !strings.Contains(text, " ") {
		return strings.Title(text)
	}

	words := strings.SplitN(text, " ", 2)

	return strings.Join([]string{
		strings.Title(words[0]),
		words[1],
	}, " ")
}

func findInt(text string, compiledRx *regexp.Regexp) (result int, found bool) {
	r, err := strconv.Atoi(rx.Spaces.ReplaceAllString(compiledRx.FindString(text), ""))
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
			hrefAttr = strings.TrimSpace(h)
			return
		}
	}
	return
}

func rawPhoneToValidPhone(in string) (phone int, err error) {
	errNotPhone := errors.New("not phone")

	numChunks := rx.Nums.FindAllString(in, -1)
	if numChunks == nil {
		err = errNotPhone
		return
	}

	nums := strings.Join(numChunks, "")
	if len(nums) != 11 {
		err = errNotPhone
		return
	}

	if string(nums[0]) == "8" {
		nums = strings.Join([]string{"7", nums[1:]}, "")
	}

	return strconv.Atoi(nums)
}

func toOGImage(imgSrc string, url string) link {
	parsedImgSrcURL, err := u.Parse(imgSrc)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return ""
	}

	baseURL, err := u.Parse(url)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return ""
	}

	if parsedImgSrcURL.Scheme == "" {
		parsedImgSrcURL.Scheme = http
	}
	if parsedImgSrcURL.Host == "" {
		punycodeHost, err := idna.New().ToASCII(baseURL.Host)
		if err != nil {
			logger.Log.Error().Err(err).Send()
			return ""
		}
		parsedImgSrcURL.Host = punycodeHost
	}

	return link(parsedImgSrcURL.String())
}

func emailSuffixValid(email string) (valid bool) {
	return !strings.HasSuffix(email, ".png")
}

func makeBool(b bool) *bool {
	return &b
}
