package vk

import (
	vkSdk "github.com/SevereCloud/vksdk/api"
	"github.com/nnqq/scr-parser/config"
	"strings"
)

var Api *vkSdk.VK

func init() {
	Api = vkSdk.NewVKWithPool(strings.Split(config.Env.Vk.Tokens, ",")...)
}
