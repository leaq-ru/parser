package company

import (
	"context"
	"net/url"
	"strings"

	"github.com/nnqq/scr-parser/logger"
)

func (c *Company) DigVk(ctx context.Context, vkUrl string) {
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

	// TODO execute = GetGroup(slug string) (GroupExecuteRes, error)
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
