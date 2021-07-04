package vk

import (
	"fmt"

	"github.com/SevereCloud/vksdk/v2/api"
)

type Service struct {
	vk *api.VK
}

func NewService(tokens []string) (*Service, error) {
	vk := api.NewVK(tokens...)
	vk.Limit = api.LimitUserToken

	// ping
	_, err := vk.AccountGetInfo(api.Params{})
	if err != nil {
		return nil, err
	}

	return &Service{
		vk: vk,
	}, nil
}

func (s *Service) Client() *api.VK {
	return s.vk
}

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
type GroupExecuteRes struct {
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
	Contacts []Contact `json:"contacts"`
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

type Contact struct {
	ID        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	IsClosed  bool    `json:"is_closed"`
	Sex       float64 `json:"sex"`
	Photo200  string  `json:"photo_200"`
}

func (s *Service) GetGroup(slug string) (GroupExecuteRes, error) {
	var res GroupExecuteRes
	err := s.vk.Execute(fmt.Sprintf(`
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
		`, slug), &res)
	if err != nil {
		return GroupExecuteRes{}, err
	}

	return res, nil
}
