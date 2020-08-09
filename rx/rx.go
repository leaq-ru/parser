package rx

import (
	"regexp"
)

var (
	Spaces = regexp.MustCompile("\\s")
	Email  = regexp.MustCompile("[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,6}")
	Nums   = regexp.MustCompile("[0-9]+")
	INN    = regexp.MustCompile("\\s[0-9]{10}\\s")
	KPP    = regexp.MustCompile("\\s[0-9]{9}\\s")
	OGRN   = regexp.MustCompile("\\s[0-9]{13}\\s")
	Phone  = regexp.MustCompile(
		"(>|\\s)(\\+7|7|8)[\\s\\-]?\\(?[489][0-9]{2}\\)?[\\s\\-]?[0-9]{3}[\\s\\-]?[0-9]{2}[\\s\\-]?[0-9]{2}")
)
