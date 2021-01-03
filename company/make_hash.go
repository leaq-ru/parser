package company

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func (c *Company) WithHash() {
	var (
		appStoreURL   string
		googlePlayURL string
	)
	if c.App != nil {
		if c.App.AppStore != nil {
			appStoreURL = c.App.AppStore.URL
		}
		if c.App.GooglePlay != nil {
			googlePlayURL = c.App.GooglePlay.URL
		}
	}

	var (
		instagramURL string
		twitterURL   string
		youtubeURL   string
		facebookURL  string
	)
	if c.GetSocial() != nil {
		if c.GetSocial().Instagram != nil {
			instagramURL = c.GetSocial().Instagram.URL
		}
		if c.GetSocial().Twitter != nil {
			twitterURL = c.GetSocial().Twitter.URL
		}
		if c.GetSocial().Youtube != nil {
			youtubeURL = c.GetSocial().Youtube.URL
		}
		if c.GetSocial().Facebook != nil {
			facebookURL = c.GetSocial().Facebook.URL
		}
	}

	sum := md5.Sum([]byte(strings.Join([]string{
		c.Title,
		c.Description,
		strconv.Itoa(c.Phone),
		c.Email,
		strconv.Itoa(c.GetSocial().GetVk().GetGroupId()),
		strconv.Itoa(c.INN),
		strconv.Itoa(c.KPP),
		strconv.Itoa(c.OGRN),
		appStoreURL,
		googlePlayURL,
		instagramURL,
		twitterURL,
		youtubeURL,
		facebookURL,
	}, ":")))

	c.Hash = hex.EncodeToString(sum[:])
}
