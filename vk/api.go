package vk

import (
	"github.com/SevereCloud/vksdk/api"
	"github.com/nnqq/scr-parser/config"
	"strings"
)

var UserApi *api.VK

func init() {
	UserApi = api.NewVKWithPool(strings.Split(config.Env.Vk.UserTokens, ",")...)
	UserApi.Limit = api.LimitUserToken
}
