package model

import (
	"context"
	"fmt"
	"github.com/nnqq/scr-parser/logger"
	"github.com/nnqq/scr-parser/vk"
	"net/url"
	"strings"
)

//	"group": {
//		"id": 144090016,
//		"name": "Каркасные авточехлы dress4car | +7 904 0555 202",
//		"screen_name": "dress4car",
//		"is_closed": 0,
//		"type": "page",
//		"is_admin": 1,
//		"admin_level": 3,
//		"is_member": 0,
//		"is_advertiser": 1,
//		"addresses": {
//			"is_enabled": true,
//			"main_address_id": 1784
//		},
//		"description": "Пошив и установка авточехлов из экокожи...",
//		"members_count": 37026,
//		"contacts": [{
//			"user_id": 421825761,
//			"desc": "Консультация и заказ",
//			"phone": "+7 904 0555 202"
//		}],
//		"photo_50": "https://sun1-25.u...NyrDcrl2Q.jpg?ava=1",
//		"photo_100": "https://sun1-23.u...do0zcQzLo.jpg?ava=1",
//		"photo_200": "https://sun1-18.u...CcGj_8RgM.jpg?ava=1"
//	},
//	"contacts": [{
//		"id": 421825761,
//		"first_name": "Андрей",
//		"last_name": "Аверьянов",
//		"is_closed": false,
//		"can_access_closed": true,
//		"sex": 2,
//		"photo_200": "https://sun1-83.u...BLFe0d6k4.jpg?ava=1"
//	}],
//	"addr": {
//		"id": 1784,
//		"address": "ул.Дачная, 1-А",
//		"city_id": 95,
//		"title": "Детейлинг центр AutoDOL"
//	},
//	"city": {
//		"id": 95,
//		"title": "Нижний Новгород"
//	}
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

func (c *Company) digVk(ctx context.Context, vkUrl string) {
	u, err := url.Parse(vkUrl)
	if err != nil {
		logger.Log.Error().Err(err).Send()
		return
	}

	groupSlug := strings.TrimPrefix(u.Path, "/")
	if groupSlug == "" {
		logger.Log.Debug().Str("vkUrl", vkUrl).Msg("group slug is empty")
		return
	}

	execute := vkExecuteRes{}
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
	err = vk.UserApi.Execute(code, &execute)
	if err != nil {
		logger.Log.Debug().Str("groupSlug", groupSlug).Msg("execute error")
		return
	}

	if execute.City.Title != "" {
		c.setCityID(ctx, strings.Join([]string{
			"г.",
			execute.City.Title,
		}, " "))
	}

	if execute.Addr.Address != "" {
		if c.Location == nil {
			c.Location = &location{}
		}
		c.Location.Address = capitalize(execute.Addr.Address)
	}
	if execute.Addr.Title != "" {
		if c.Location == nil {
			c.Location = &location{}
		}
		c.Location.AddressTitle = capitalize(execute.Addr.Title)
	}

	userMoreFields := map[float64]vkExecuteContact{}
	for _, contact := range execute.Contacts {
		userMoreFields[contact.ID] = contact
	}

	if len(c.People) == 0 {
		for _, contact := range execute.Group.Contacts {
			item := peopleItem{
				VkID:        int(contact.UserID),
				Email:       strings.TrimSpace(contact.Email),
				Description: capitalize(contact.Desc),
			}

			user, ok := userMoreFields[contact.UserID]
			if ok {
				item.FirstName = capitalize(user.FirstName)
				item.LastName = capitalize(user.LastName)
				item.VkIsClosed = user.IsClosed
				item.Sex = int8(user.Sex)
				item.Photo200 = link(user.Photo200)
			}

			phone, err := rawPhoneToValidPhone(contact.Phone)
			if err == nil {
				item.Phone = phone
			}

			c.People = append(c.People, &item)
		}
	}

	if c.Social == nil {
		c.Social = &social{}
	}
	if c.Social.Vk == nil {
		c.Social.Vk = &vkItem{}
	}
	c.Social.Vk.GroupID = int(execute.Group.ID)
	c.Social.Vk.Name = capitalize(execute.Group.Name)
	c.Social.Vk.ScreenName = execute.Group.ScreenName
	c.Social.Vk.IsClosed = int8(execute.Group.IsClosed)
	c.Social.Vk.Description = capitalize(execute.Group.Description)
	c.Social.Vk.MembersCount = int(execute.Group.MembersCount)
	c.Social.Vk.Photo200 = link(execute.Group.Photo200)

	return
}
