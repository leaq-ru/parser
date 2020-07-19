package company

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/nnqq/scr-parser/rx"
	"golang.org/x/net/html/charset"
	"io/ioutil"
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
	numChunks := rx.Nums.FindAllString(in, -1)
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

func convertToUTF8(in []byte, origEncoding string) (res []byte, err error) {
	byteReader := bytes.NewReader(in)
	reader, err := charset.NewReaderLabel(origEncoding, byteReader)
	if err != nil {
		return
	}
	return ioutil.ReadAll(reader)
}
